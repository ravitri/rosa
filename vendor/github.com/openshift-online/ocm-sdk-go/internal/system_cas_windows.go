//go:build windows
// +build windows

/*
Copyright (c) 2021 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file contains the function that returns the trusted CA certificates for Windows. This is
// needed because currently Go doesn't know how to load the Windows trusted CA store. See the
// following issues for more information:
//
//	https://github.com/golang/go/issues/16736
//	https://github.com/golang/go/issues/18609

package internal

import (
	"crypto/x509"
)

// loadSystemCAs loads the certificates of the CAs that we will trust. Currently this uses a fixed
// set of CA certificates, which is obviusly going to break in the future, but there is not much we
// can do (or know to do) till Go learns to read the Windows CA trust store.
func loadSystemCAs() (pool *x509.CertPool, err error) {
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(ssoCA1)
	pool.AppendCertsFromPEM(ssoCA2)
	pool.AppendCertsFromPEM(apiCA1)
	pool.AppendCertsFromPEM(apiCA2)
	return
}

// The SSO certificates has been obtained with the following command:
//
//	$ openssl s_client -connect sso.redhat.com:443 -showcerts
var ssoCA1 = []byte(`
-----BEGIN CERTIFICATE-----
MIIGaDCCBVCgAwIBAgIQC46dyjYn3u1vs/8YgztPbTANBgkqhkiG9w0BAQsFADB1
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMTQwMgYDVQQDEytEaWdpQ2VydCBTSEEyIEV4dGVuZGVk
IFZhbGlkYXRpb24gU2VydmVyIENBMB4XDTIzMDcxMzAwMDAwMFoXDTI0MDcxMjIz
NTk1OVowgcoxEzARBgsrBgEEAYI3PAIBAxMCVVMxGTAXBgsrBgEEAYI3PAIBAhMI
RGVsYXdhcmUxHTAbBgNVBA8MFFByaXZhdGUgT3JnYW5pemF0aW9uMRAwDgYDVQQF
EwcyOTQ1NDM2MQswCQYDVQQGEwJVUzEXMBUGA1UECBMOTm9ydGggQ2Fyb2xpbmEx
EDAOBgNVBAcTB1JhbGVpZ2gxFjAUBgNVBAoTDVJlZCBIYXQsIEluYy4xFzAVBgNV
BAMTDnNzby5yZWRoYXQuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYrJM
qHET7N7dhi5EMpERt3tW+jC1VkZZ/CsayqoQeB2Sfzwrwe3e/j6eShNqc6XN0Z68
3B18Kfar0zWSK5CwvKOCA2cwggNjMB8GA1UdIwQYMBaAFD3TUKXWoK3u80pgCmXT
IdT4+NYPMB0GA1UdDgQWBBTx9cmugNxKMPgWsjc/6UoqfQK/nDAZBgNVHREEEjAQ
gg5zc28ucmVkaGF0LmNvbTAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0lBBYwFAYIKwYB
BQUHAwEGCCsGAQUFBwMCMHUGA1UdHwRuMGwwNKAyoDCGLmh0dHA6Ly9jcmwzLmRp
Z2ljZXJ0LmNvbS9zaGEyLWV2LXNlcnZlci1nMy5jcmwwNKAyoDCGLmh0dHA6Ly9j
cmw0LmRpZ2ljZXJ0LmNvbS9zaGEyLWV2LXNlcnZlci1nMy5jcmwwSgYDVR0gBEMw
QTALBglghkgBhv1sAgEwMgYFZ4EMAQEwKTAnBggrBgEFBQcCARYbaHR0cDovL3d3
dy5kaWdpY2VydC5jb20vQ1BTMIGIBggrBgEFBQcBAQR8MHowJAYIKwYBBQUHMAGG
GGh0dHA6Ly9vY3NwLmRpZ2ljZXJ0LmNvbTBSBggrBgEFBQcwAoZGaHR0cDovL2Nh
Y2VydHMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0U0hBMkV4dGVuZGVkVmFsaWRhdGlv
blNlcnZlckNBLmNydDAJBgNVHRMEAjAAMIIBfAYKKwYBBAHWeQIEAgSCAWwEggFo
AWYAdgDuzdBk1dsazsVct520zROiModGfLzs3sNRSFlGcR+1mwAAAYlQB49zAAAE
AwBHMEUCIGlDT2c7z2fv71iyJbuNYbp91KGp0Koy6qr3So9K4YiGAiEA80VIq3yN
Vu9WchFtK0xQh8pfMqGefkwGrJNWhb9B2AsAdQBIsONr2qZHNA/lagL6nTDrHFIB
y1bdLIHZu7+rOdiEcwAAAYlQB4/RAAAEAwBGMEQCIEG+h90k/erQqUhsFk6guFOT
zfm5TZ9JngeUayQ6lnoPAiBSYCHSjVyt5CUMOuJzwJ7bH0Cszfj6hXb3E9Vv4vkF
xgB1ANq2v2s/tbYin5vCu1xr6HCRcWy7UYSFNL2kPTBI1/urAAABiVAHj6QAAAQD
AEYwRAIgctt1pUVJPnAEA4c72rcbRKXZ+IMzZSBjv//hOtst1eYCIDeEu5I0kdaF
bz6EucRcalZfMDtnwn+FMwkgRbiqo6l2MA0GCSqGSIb3DQEBCwUAA4IBAQA+mnI3
mcDI1krs4JYzlHAzPs5IgP6PVLeBsdPq2AJBAHLmQsYPBKPWjJi744O+kmfwGH3M
s7ZplC2GgYBW3izfcXcf6aw0UrfVizRz8MNNTknjHwP21FwS1ykpjlJL0QvGM67g
3koAne4fsHmSOgFNchaPnm3QbKQdb9oXMX+HPY/xxwvwQeUWs94g0C5/lHM8+7Tq
GvtAzh/3EdWN5ZwOMHXE9nGvnpM7Aj9A9nuMvwMsn6YqeBaq71UirB7VELUyjSVc
rVl7msAPTJd8vR+NtY4fTSanNnQDptG2M1OcPthH/rwINE0rdHaVFDDRSz4G78f7
l1X92ab6Xau+LI4u
-----END CERTIFICATE-----
`)

var ssoCA2 = []byte(`
-----BEGIN CERTIFICATE-----
MIIEtjCCA56gAwIBAgIQDHmpRLCMEZUgkmFf4msdgzANBgkqhkiG9w0BAQsFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTEzMTAyMjEyMDAwMFoXDTI4MTAyMjEyMDAwMFowdTEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTE0MDIGA1UEAxMrRGlnaUNlcnQgU0hBMiBFeHRlbmRlZCBW
YWxpZGF0aW9uIFNlcnZlciBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBANdTpARR+JmmFkhLZyeqk0nQOe0MsLAAh/FnKIaFjI5j2ryxQDji0/XspQUY
uD0+xZkXMuwYjPrxDKZkIYXLBxA0sFKIKx9om9KxjxKws9LniB8f7zh3VFNfgHk/
LhqqqB5LKw2rt2O5Nbd9FLxZS99RStKh4gzikIKHaq7q12TWmFXo/a8aUGxUvBHy
/Urynbt/DvTVvo4WiRJV2MBxNO723C3sxIclho3YIeSwTQyJ3DkmF93215SF2AQh
cJ1vb/9cuhnhRctWVyh+HA1BV6q3uCe7seT6Ku8hI3UarS2bhjWMnHe1c63YlC3k
8wyd7sFOYn4XwHGeLN7x+RAoGTMCAwEAAaOCAUkwggFFMBIGA1UdEwEB/wQIMAYB
Af8CAQAwDgYDVR0PAQH/BAQDAgGGMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEF
BQcDAjA0BggrBgEFBQcBAQQoMCYwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRp
Z2ljZXJ0LmNvbTBLBgNVHR8ERDBCMECgPqA8hjpodHRwOi8vY3JsNC5kaWdpY2Vy
dC5jb20vRGlnaUNlcnRIaWdoQXNzdXJhbmNlRVZSb290Q0EuY3JsMD0GA1UdIAQ2
MDQwMgYEVR0gADAqMCgGCCsGAQUFBwIBFhxodHRwczovL3d3dy5kaWdpY2VydC5j
b20vQ1BTMB0GA1UdDgQWBBQ901Cl1qCt7vNKYApl0yHU+PjWDzAfBgNVHSMEGDAW
gBSxPsNpA/i/RwHUmCYaCALvY2QrwzANBgkqhkiG9w0BAQsFAAOCAQEAnbbQkIbh
hgLtxaDwNBx0wY12zIYKqPBKikLWP8ipTa18CK3mtlC4ohpNiAexKSHc59rGPCHg
4xFJcKx6HQGkyhE6V6t9VypAdP3THYUYUN9XR3WhfVUgLkc3UHKMf4Ib0mKPLQNa
2sPIoc4sUqIAY+tzunHISScjl2SFnjgOrWNoPLpSgVh5oywM395t6zHyuqB8bPEs
1OG9d4Q3A84ytciagRpKkk47RpqF/oOi+Z6Mo8wNXrM9zwR4jxQUezKcxwCmXMS1
oVWNWlZopCJwqjyBcdmdqEU79OX2olHdx3ti6G8MdOu42vi/hw15UJGQmxg7kVkn
8TUoE6smftX3eg==
-----END CERTIFICATE-----
`)

// The API certificates have been obtained with the following command:
//
//	$ openssl s_client -connect api.openshift.com:443 -showcerts
var apiCA1 = []byte(`
-----BEGIN CERTIFICATE-----
MIIE7zCCA9egAwIBAgISBBfcQDDi+eK9uvPWwZSjTU3rMA0GCSqGSIb3DQEBCwUA
MDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD
EwJSMzAeFw0yNDA0MjIxMjA4MTRaFw0yNDA3MjExMjA4MTNaMBwxGjAYBgNVBAMT
EWFwaS5vcGVuc2hpZnQuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC
AQEAqRG7wobhclyQJp7jbr3+OzTYot1RiPgFXKjQe2K3CLxcx1fA5BTdw6tCv7wr
3Atkgi68S/jsmpdeFb4c4caD5IUNZpiD977wWKcsCVXkXcIttQf0POTyauna3Amu
3zRqVXA6vYYOK+qHpBM+iO2u1vtTwRhooxXDgRtCWQdcb7U8sdttFT9Bk4tny6a1
ioQQH3MnLsnWc3qIzh1aoiSr6UVJdGDrdiIfpDvpp/nbZilprGBhM3MXHDI1rLyo
NlVzxKQnxIPPO0gNwPFyWVsF3mQ1mgPICWSBVJ3pyVSe9iAvpwNHO+xCOxIM9oST
gcRnbPIHr0QnXCoRw6Vj2YlJ/wIDAQABo4ICEzCCAg8wDgYDVR0PAQH/BAQDAgWg
MB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0G
A1UdDgQWBBSIbbXrfs+0652bZEFCfZbiAP2sFTAfBgNVHSMEGDAWgBQULrMXt1hW
y65QCUDmH6+dixTCxjBVBggrBgEFBQcBAQRJMEcwIQYIKwYBBQUHMAGGFWh0dHA6
Ly9yMy5vLmxlbmNyLm9yZzAiBggrBgEFBQcwAoYWaHR0cDovL3IzLmkubGVuY3Iu
b3JnLzAcBgNVHREEFTATghFhcGkub3BlbnNoaWZ0LmNvbTATBgNVHSAEDDAKMAgG
BmeBDAECATCCAQQGCisGAQQB1nkCBAIEgfUEgfIA8AB2AD8XS0/XIkdYlB1lHIS+
DRLtkDd/H4Vq68G/KIXs+GRuAAABjwXr3kgAAAQDAEcwRQIhANiMWi6rqpsslXxM
qC0Txq8kAuzO2zXTInIDB1ET6tI5AiBizAQcIFrKd2AiHJCOSVVVv0zrMb5RtOyv
GSKzBn5fHgB2ABmYEHEJ8NZSLjCA0p4/ZLuDbijM+Q9Sju7fzko/FrTKAAABjwXr
3msAAAQDAEcwRQIgHUdPSu8gY3QzmM+47/4PLANq9COBEqUR7qJ4YMJ5380CIQC/
8nKMoLod4WrvAQI7JxBWN+H9dwd7JuzmB0Wt0qiTzDANBgkqhkiG9w0BAQsFAAOC
AQEACc/kolajlphiHPuRUr62RNlurGJjKqaaHXegH2yEPCawQnd93+8LOnf/udV3
zK2XDQRtX/n3Y1eULogkHZMDpzrz/9ft46s258hqIoYVNzYstb8EG60oOiA01rYv
Pe1+KniG481J00I8OK4emQz85H6cEHOth6MHJtQNxUerpF/m/712uOiNwld3GQRG
Uh04EElLhnbnpNa7Ki8Xf7sRFrUU1M0ijiNqf8L0RHOXOthCfvjCdzpZ0gREfZFS
uyS5Q5j8G30Of3DOsLgzCSjfKgHBml4jeLuxtdQopea8FmpqGhEryLDJmR+jNxDH
eW8SsdjjvIKA33jrZIKYhwEUnQ==
-----END CERTIFICATE-----
`)

var apiCA2 = []byte(`
-----BEGIN CERTIFICATE-----
MIIFFjCCAv6gAwIBAgIRAJErCErPDBinU/bWLiWnX1owDQYJKoZIhvcNAQELBQAw
TzELMAkGA1UEBhMCVVMxKTAnBgNVBAoTIEludGVybmV0IFNlY3VyaXR5IFJlc2Vh
cmNoIEdyb3VwMRUwEwYDVQQDEwxJU1JHIFJvb3QgWDEwHhcNMjAwOTA0MDAwMDAw
WhcNMjUwOTE1MTYwMDAwWjAyMQswCQYDVQQGEwJVUzEWMBQGA1UEChMNTGV0J3Mg
RW5jcnlwdDELMAkGA1UEAxMCUjMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQC7AhUozPaglNMPEuyNVZLD+ILxmaZ6QoinXSaqtSu5xUyxr45r+XXIo9cP
R5QUVTVXjJ6oojkZ9YI8QqlObvU7wy7bjcCwXPNZOOftz2nwWgsbvsCUJCWH+jdx
sxPnHKzhm+/b5DtFUkWWqcFTzjTIUu61ru2P3mBw4qVUq7ZtDpelQDRrK9O8Zutm
NHz6a4uPVymZ+DAXXbpyb/uBxa3Shlg9F8fnCbvxK/eG3MHacV3URuPMrSXBiLxg
Z3Vms/EY96Jc5lP/Ooi2R6X/ExjqmAl3P51T+c8B5fWmcBcUr2Ok/5mzk53cU6cG
/kiFHaFpriV1uxPMUgP17VGhi9sVAgMBAAGjggEIMIIBBDAOBgNVHQ8BAf8EBAMC
AYYwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBIGA1UdEwEB/wQIMAYB
Af8CAQAwHQYDVR0OBBYEFBQusxe3WFbLrlAJQOYfr52LFMLGMB8GA1UdIwQYMBaA
FHm0WeZ7tuXkAXOACIjIGlj26ZtuMDIGCCsGAQUFBwEBBCYwJDAiBggrBgEFBQcw
AoYWaHR0cDovL3gxLmkubGVuY3Iub3JnLzAnBgNVHR8EIDAeMBygGqAYhhZodHRw
Oi8veDEuYy5sZW5jci5vcmcvMCIGA1UdIAQbMBkwCAYGZ4EMAQIBMA0GCysGAQQB
gt8TAQEBMA0GCSqGSIb3DQEBCwUAA4ICAQCFyk5HPqP3hUSFvNVneLKYY611TR6W
PTNlclQtgaDqw+34IL9fzLdwALduO/ZelN7kIJ+m74uyA+eitRY8kc607TkC53wl
ikfmZW4/RvTZ8M6UK+5UzhK8jCdLuMGYL6KvzXGRSgi3yLgjewQtCPkIVz6D2QQz
CkcheAmCJ8MqyJu5zlzyZMjAvnnAT45tRAxekrsu94sQ4egdRCnbWSDtY7kh+BIm
lJNXoB1lBMEKIq4QDUOXoRgffuDghje1WrG9ML+Hbisq/yFOGwXD9RiX8F6sw6W4
avAuvDszue5L3sz85K+EC4Y/wFVDNvZo4TYXao6Z0f+lQKc0t8DQYzk1OXVu8rp2
yJMC6alLbBfODALZvYH7n7do1AZls4I9d1P4jnkDrQoxB3UqQ9hVl3LEKQ73xF1O
yK5GhDDX8oVfGKF5u+decIsH4YaTw7mP3GFxJSqv3+0lUFJoi5Lc5da149p90Ids
hCExroL1+7mryIkXPeFM5TgO9r0rvZaBFOvV2z0gp35Z0+L4WPlbuEjN/lxPFin+
HlUjr8gRsI3qfJOQFy/9rKIJR0Y/8Omwt/8oTWgy1mdeHmmjk7j1nYsvC9JSQ6Zv
MldlTTKB3zhThV1+XWYp6rjd5JW1zbVWEkLNxE7GJThEUG3szgBVGP7pSWTUTsqX
nLRbwHOoq7hHwg==
-----END CERTIFICATE-----
`)
