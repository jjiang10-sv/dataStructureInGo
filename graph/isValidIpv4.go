package graph

// import (
// 	"regexp"
// 	"fmt"
// )

// // Regular expression pattern for IPv4 validation
// var ipv4Pattern = `^(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.` +
//                   `(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.` +
//                   `(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.` +
//                   `(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])$`

// func isValidIPv4(ip string) bool {
//     re := regexp.MustCompile(ipv4Pattern)
//     return re.MatchString(ip)
// }

// func mainValid() {
//     testIPs := []string{
//         "192.168.1.1",
//         "255.255.255.255",
//         "0.0.0.0",
//         "256.100.50.0",     // Invalid
//         "192.168.1.300",    // Invalid
//         "123.045.067.089",  // Invalid
//     }

//     for _, ip := range testIPs {
//         if isValidIPv4(ip) {
//             fmt.Printf("%s is a valid IPv4 address.\n", ip)
//         } else {
//             fmt.Printf("%s is NOT a valid IPv4 address.\n", ip)
//         }
//     }
// }
