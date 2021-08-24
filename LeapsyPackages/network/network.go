package network

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"leapsy.com/packages/configurations"
	"leapsy.com/packages/logings"
)

var (
	addressToAliasMap = make(map[string]string) // 網址對應別名
	readWriteLock     = new(sync.RWMutex)       // 讀寫鎖
)

// LookupHostString - 查找主機
/**
 * @param  string hostString  主機
 * @return []string results 查找主機結果
 */
func LookupHostString(hostString string) (results []string) {

	// 主機正規表示式
	hostStringRegularExpression :=
		`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9]):\d{1,5}$`

	if regexp.MustCompile(hostStringRegularExpression).MatchString(hostString) { // 若主機符合格式

		hostSlices := strings.Split(hostString, `:`) // 將主機名稱切開

		formatStringSlices := []string{`查找主機 %s `} // 記錄器格式片段
		defaultArgs := []interface{}{hostString}   // 記錄器預設參數

		addrs, netLookupHostError := net.LookupHost(hostSlices[0]) // 查找主機

		logings.SendLog(
			formatStringSlices,
			defaultArgs,
			netLookupHostError,
			0,
		)

		if nil != netLookupHostError { // 若查找主機錯誤
			return // 回傳
		}

		for _, addr := range addrs { // 針對每一結果

			updatedHostString := addr + `:` + hostSlices[1] // 結果加埠

			if _, ok := addressToAliasMap[updatedHostString]; ok { // 若結果加埠存在別名
				results = append(results, updatedHostString) // 將結果加埠加入回傳
			}

		}

		logings.SendLog(
			[]string{`取得主機 %s 資料 %+v `},
			append(defaultArgs, results),
			nil,
			0,
		)

	} else { // 若主機不符合格式，則記錄資訊

		logings.SendLog(
			[]string{`主機查找 %s `},
			[]interface{}{hostString},
			fmt.Errorf(`主機不符合格式 [host]:[IP]`),
			logrus.InfoLevel,
		)

	}

	return // 回傳
}

// GetAddressAlias - 取得位址別名
/**
 * @param  string addressString 位址字串
 * @return string 位址別名
 */
func GetAddressAlias(addressString string) string {
	readWriteLock.RLock()                   // 讀鎖
	defer readWriteLock.RUnlock()           // 記得解開讀鎖
	return addressToAliasMap[addressString] // 回傳位址別名
}

// SetAddressAlias - 設定位址別名
/**
 * @param  string addressString 位址字串
 * @param  string aliasString 別名字串
 */
func SetAddressAlias(addressString, aliasString string) {
	readWriteLock.Lock()                           // 寫鎖
	defer readWriteLock.Unlock()                   // 記得解開寫鎖
	addressToAliasMap[addressString] = aliasString // 設定位址別名
}

// GetAliasAddressPair - 取得(別名,位址)對
/**
 * @param  string addressString 位址字串
 * @return  []interface{} (別名,位址)對
 */
func GetAliasAddressPair(addressString string) []interface{} {
	return []interface{}{GetAddressAlias(addressString), addressString} // 回傳(別名,位址)對
}

func isIPv6GlobalUnicast(address net.IP) bool {

	globalUnicastIPv6Net := net.IPNet{
		IP:   net.IP{0x20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(3, 128),
	}

	return globalUnicastIPv6Net.Contains(address)
}

// mustParseCIDR - parses string into net.IPNet
/*
 * @param inputString string 輸入字串
 * @return ipNet net.IPNet 回傳結果
 */
func mustParseCIDR(inputString string) (ipNet net.IPNet) {

	_, ipNetPointer, parseCIDRError := net.ParseCIDR(inputString)

	logings.SendLog(
		[]string{`解析 '%s' 為 net.IPNet指標 '%+v' `},
		[]interface{}{inputString, ipNetPointer},
		parseCIDRError,
		0,
	)

	if nil != parseCIDRError { // 若解析錯誤
		return // 回傳
	}

	ipNet = *ipNetPointer

	return

}

// isIPv4Reserved - 判斷是否為保留IP
/*
 * @param net.IP address 網址
 * @return bool 結果
 */
func isIPv4Reserved(address net.IP) bool {

	// reservedIPv4Nets net.IPNets that are loopback, private, link local, default unicast
	// based on https://github.com/letsencrypt/boulder/blob/master/bdns/dns.go
	reservedIPv4Nets := []net.IPNet{
		mustParseCIDR("10.0.0.0/8"),         // RFC1918
		mustParseCIDR("172.16.0.0/12"),      // private
		mustParseCIDR("192.168.0.0/16"),     // private
		mustParseCIDR("127.0.0.0/8"),        // RFC5735
		mustParseCIDR("0.0.0.0/8"),          // RFC1122 Section 3.2.1.3
		mustParseCIDR("169.254.0.0/16"),     // RFC3927
		mustParseCIDR("192.0.0.0/24"),       // RFC 5736
		mustParseCIDR("192.0.2.0/24"),       // RFC 5737
		mustParseCIDR("198.51.100.0/24"),    // Assigned as TEST-NET-2
		mustParseCIDR("203.0.113.0/24"),     // Assigned as TEST-NET-3
		mustParseCIDR("192.88.99.0/24"),     // RFC 3068
		mustParseCIDR("192.18.0.0/15"),      // RFC 2544
		mustParseCIDR("224.0.0.0/4"),        // RFC 3171
		mustParseCIDR("240.0.0.0/4"),        // RFC 1112
		mustParseCIDR("255.255.255.255/32"), // RFC 919 Section 7
		mustParseCIDR("100.64.0.0/10"),      // RFC 6598
		mustParseCIDR("::/128"),             // RFC 4291: Unspecified Address
		mustParseCIDR("::1/128"),            // RFC 4291: Loopback Address
		mustParseCIDR("100::/64"),           // RFC 6666: Discard Address Block
		mustParseCIDR("2001::/23"),          // RFC 2928: IETF Protocol Assignments
		mustParseCIDR("2001:2::/48"),        // RFC 5180: Benchmarking
		mustParseCIDR("2001:db8::/32"),      // RFC 3849: Documentation
		mustParseCIDR("2001::/32"),          // RFC 4380: TEREDO
		mustParseCIDR("fc00::/7"),           // RFC 4193: Unique-Local
		mustParseCIDR("fe80::/10"),          // RFC 4291: Section 2.5.6 Link-Scoped Unicast
		mustParseCIDR("ff00::/8"),           // RFC 4291: Section 2.7
		mustParseCIDR("2002::/16"),          // RFC 7526: 6to4 anycast prefix deprecated
	}

	for _, reservedNet := range reservedIPv4Nets {

		if reservedNet.Contains(address) {
			return true
		}

	}

	return false
}

// isPublicIPAddress - 判斷是否 public IP
/*
 * @param net.IP address 網址
 * @return bool 結果
 */
func isPublicIPAddress(address net.IP) bool {

	if address.To4() != nil {
		return !isIPv4Reserved(address)
	}

	return isIPv6GlobalUnicast(address)
}

// isOutsideIPString - 判斷是否外部IP字串
/*
 * @param net.IP address 網址
 * @return result bool 結果
 */
func isOutsideIPString(ipString string) (result bool) {

	defaultArgs := []interface{}{} // 預設參數

	url := configurations.GetConfigValueOrPanic(`servers.ClockInAPIServer`, `company-IP-URL`)
	getResponse, getError := http.Get(url)

	logings.SendLog(
		[]string{`GET '%s' `},
		append(defaultArgs, url),
		getError,
		0,
	)

	if nil != getError { // 若GET錯誤
		return // 回傳
	}

	defer getResponse.Body.Close()

	publicIPBytes, readAllError := ioutil.ReadAll(getResponse.Body)
	publicIPString := string(publicIPBytes)

	logings.SendLog(
		[]string{`從 '%s' 取得主機公共IP '%s' `},
		append(defaultArgs, url, publicIPString),
		getError,
		0,
	)

	if nil != readAllError { // 若取得主機公共IP錯誤
		return // 回傳
	}

	iPNet, _, _ := net.ParseCIDR(ipString)
	result = isPublicIPAddress(iPNet) && publicIPString != ipString

	return // 回傳
}

// VerifyInsideIPString - 驗證內部IP字串
/*
 * @param net.IP address 網址
 * @return result error 結果
 */
func VerifyInsideIPString(ipString string) (result error) {

	if isOutsideIPString(ipString) {
		result = errors.New(ipIsOutsideErrorConstString)
	}

	return // 回傳
}
