package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	type Coverage struct {
		URI string `xml:"uri,attr"`
	}
	type External struct {
		Coverage Coverage `xml:"coverage"`
	}
	type Alert struct {
	}
	type TestMethod struct {
		URI           string  `xml:"uri,attr"`
		Type          string  `xml:"type,attr"`
		Name          string  `xml:"name,attr"`
		PackageName   string  `xml:"packageName,attr"`
		ExecutionTime int     `xml:"executionTime,attr"`
		Unit          string  `xml:"unit,attr"`
		Alerts        []Alert `xml:"alerts>alert"`
	}
	type TestClass struct {
		URI         string       `xml:"uri,attr"`
		Type        string       `xml:"type,attr"`
		Name        string       `xml:"name,attr"`
		PackageName string       `xml:"packageName,attr"`
		Alerts      []Alert      `xml:"alerts>alert"`
		TestMethods []TestMethod `xml:"testMethods>testMethod"`
	}
	type Program struct {
		URI         string      `xml:"uri,attr"`
		Type        string      `xml:"type,attr"`
		Name        string      `xml:"name,attr"`
		PackageName string      `xml:"packageName,attr"`
		Alerts      []Alert     `xml:"alerts>alert"`
		TestClasses []TestClass `xml:"testClasses>testClass"`
	}
	type RunResult struct {
		XMLName  xml.Name  `xml:"runResult"`
		External External  `xml:"external"`
		Alerts   []Alert   `xml:"alerts>alert"`
		Programs []Program `xml:"program"`
	}

	data := `<?xml version="1.0" encoding="utf-8"?><aunit:runResult xmlns:aunit="http://www.sap.com/adt/aunit"><external><coverage adtcore:uri="/sap/bc/adt/runtime/traces/coverage/measurements/0800279158BE1ED8AFFA4794B38F5A3A" xmlns:adtcore="http://www.sap.com/adt/core"/></external><alerts/><program adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext" adtcore:type="CLAS/OC" adtcore:name="ZCL_ZSADL_TEST_DPC_EXT" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"><alerts/><testClasses><testClass adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/includes/testclasses#start=12,6" adtcore:type="CLAS/OCN/testclasses" adtcore:name="LTCL_TEST" adtcore:packageName="$SADL_TEST"><alerts/><testMethods><testMethod adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/includes/testclasses#start=14,9" adtcore:type="CLAS/OCN/testclasses" adtcore:name="DUMMY" adtcore:packageName="$SADL_TEST" executionTime="0" unit="s"><alerts/></testMethod></testMethods></testClass></testClasses></program></aunit:runResult>`
	v := RunResult{}

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("External: %v\n", v.External)
	fmt.Printf("Alerts: %v\n", v.Alerts)
	fmt.Printf("Programs: %v\n", v.Programs)
}
