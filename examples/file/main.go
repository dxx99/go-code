package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	objectIds = []string{
		"5c88c7f3a10a3e2acf1a7d93",
		"5ce3a1dba10a3e7154334561",
		"5ce50b1da10a3e714b7e826d",
		"5d109119a10a3e604a5f5396",
		"5d131016ea7b19001d19f3c2",
		"5d14062bb00ebd001e5ec5c2",
		"5d14b436eecec6001b637002",
		"5d1614f34fa33a001d3e1a4b",
		"5d1dc0682d0de6001d4d5d14",
		"5d23fd30b513b8001c26e592",
		"5d3576ac12c461001b637de2",
		"5d36cb1e12c461001d2880f2",
		"5d3e94ddd0a47900180839ae",
		"5d3e96aad0a47900180839bb",
		"5d3ea7d4d0a47900180839cb",
		"5d3ec9491e41ad001809cdc2",
		"5d3ffe016b04040018499112",
		"5d3ffe016b04040018499114",
		"5d413a9a0457f200186a2f32",
		"5d413afb0457f200186a2f35",
		"5d4258c206b6a900181fb56f",
		"5d42b5b1660c7100185f4892",
		"5d43986b660c7100185f489c",
		"5d439c06660c7100185f489d",
		"5d439d40660c7100185f48a1",
		"5d439e06660c7100185f48a5",
		"5d440f8a660c7100185f48b5",
		"5d4797fa660c7100185f48c3",
		"5d47a35d660c7100185f48c7",
		"5d47a3b2660c7100185f48cb",
		"5d47c820660c7100185f48d0",
		"5d48e05252f0a600181573d8",
		"5d71f961918cb4001875d254",
		"5dd1fd9ebcfd0000182724b2",
		"5dd34bb4c01bc3001875f754",
		"5ddcc40d965aad00186ea152",
		"5ddceb90965aad00186ea16b",
		"5ddd1047965aad00186ea17a",
		"5dde46c7dc93b700186b2417",
		"5de08f0707f33c0018410e55",
		"5de9f6f7fb7bb500186824c2",
		"5eb69570bcd45500e560cee3",
		"5ec20f95cce93800e66d96e2",
		"5d7871d96d42ec00187c661e",
		"5d81db356750b00018258462",
		"5d81db6e6750b00018258466",
		"5d81db956750b0001825846a",
		"5d81dc1f6750b0001825846e",
		"5d81e2f36750b00018258472",
		"5d81e3706750b00018258476",
		"5d81e7ce6750b0001825847a",
		"5d81e9c96750b0001825847e",
		"5d833d3613e1f900180e5393",
		"5d834d6313e1f900180e5397",
		"5d8358ca13e1f900180e539b",
		"5d84902c13e1f900180e53a5",
		"5d849bd113e1f900180e53ad",
		"5d8c91ae5542a60018356656",
		"5d8d88935542a60018356665",
		"5ec20a07cce93800be5b64ed",
		"5d394f58b02537001936df84",
		"5d395b7b5df43d0018487182",
		"5d4cd8653ba5ef00180cdfbe",
		"5d4cd97b3ba5ef00180cdfc3",
		"5d4d0f913ba5ef00180cdfc8",
		"5d54bd71a7d12f00187bcc9a",
		"5d54c531a7d12f00187bcca0",
		"5d54db5114a74500180a4fa4",
		"5d662a45d60266001820add3",
		"5d9064515542a60018356672",
		"5da16ea421267300187438b2",
		"5da173d821267300187438b8",
		"5da1741721267300187438bc",
		"5da81513d132aa005e1dcca9",
		"5de0b98307f33c0018410e69",
		"5ec2057bcce93800be5b64e9",
		"5ec246a2cce9380187294fa3",
		"5f18f1d9532d3c0032570db3",
		"5f18f1e2532d3c0032570db6",
		"5f3a28bafc7290006e46d432",
		"5f3cc56a3126980018031fc9",
		"5f3de6ae312698006279f85b",
		"5f59c7b90b7b8300171e3c71",
		"5f61828f0b7b83009c539272",
		"5fed4c95d3d2de000e723b63",
		"5ff533d6d3d2de000e723b78",
		"60b778706c3fbf00167322b1",
		"60bd8bd0ed5589001529fe7d",
		"60c1944ba32fd6001672e40e",
		"60c194b1a32fd6001565ec97",
		"60c1b2fca32fd6001565ec9b",
		"60c1b409a32fd6001757427c",
		"60c1b4b7a32fd60017574280",
		"60c1b4e9a32fd60017574284",
		"60c1bcb9a32fd600184ea334",
		"60c305d9a32fd6001672e413",
		"610cfca97c808b00162602da",
		"611b6d2ce65c4e00195419c7",
		"611b8118545d49001166f8c3",
		"611c94a989ff6e0016592c24",
		"6177b3c8bd065b001803487a",
		"619e06f9de7f250011300542",
		"61a84af25110070f03386032",
		"61d6ab767cc97b08a84f51b2",
		"61d6ab817cc97b0c00616f02",
		"628456e9db2ea400172a08c3",
		"62845706db2ea4001535aba2",
		"62878e6c2ad47f0015176a57"}
)

type storePoint struct {
	ObjectId string
	Detail string
	Lat string
	Lng string
}

func main() {
	dir, err := os.Getwd()
	f, err := os.Open(dir + "/file/1.txt")
	if err != nil {
		panic(err)
	}
	//
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fs, err := os.OpenFile(dir + "/file/output.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	res  := strings.Split(string(bs), "\n")
	k := 0
	fmt.Println(len(objectIds)-1)
	for _, item := range res {
		if k > len(objectIds)-1 {
			break
		}
		regLat := regexp.MustCompile(`""lat"" : ([\s\S]+),`)
		lat := regLat.FindAllStringSubmatch(item, -1)
		if len(lat) == 0 || len(lat[0][1]) < 8 {
			continue
		}

		regLng := regexp.MustCompile(`""lng"" : ([\s\S]+) }`)
		lng := regLng.FindAllStringSubmatch(item, -1)
		if len(lng) == 0 || len(lng[0][1]) < 8 {
			continue
		}

		regDetail := regexp.MustCompile(`""detail"" : ""([\s\S]+)"" }`)
		detail := regDetail.FindAllStringSubmatch(item, -1)
		if len(detail) == 0 || len(detail[0][1]) < 5 {
			continue
		}


		s := fmt.Sprintf(`db.getCollection("saler").update({"_id":ObjectId("%s")}, {"$set":{"store_address.detail":"%s", "loc_point.lng":%s, "loc_point.lat":%s}});`,
			objectIds[k], detail[0][1], lng[0][1], lat[0][1])
		fmt.Println(s)
		fs.WriteString(s + "\n")
		k++
	}
	fs.Close()


}
