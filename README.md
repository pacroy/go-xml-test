# go-xml-test
Test Go encoding/xml

# ABAPUnit

## Request Body

```xml
<?xml version="1.0" encoding="UTF-8"?>
<aunit:runConfiguration xmlns:aunit="http://www.sap.com/adt/aunit">
  <external>
    <coverage active="true"/>
  </external>
  <adtcore:objectSets xmlns:adtcore="http://www.sap.com/adt/core">
    <objectSet kind="inclusive">
      <adtcore:objectReferences>
		<adtcore:objectReference adtcore:uri="/sap/bc/adt/vit/wb/object_type/devck/object_name/$SADL_TEST"/>
      </adtcore:objectReferences>
    </objectSet>
  </adtcore:objectSets>
</aunit:runConfiguration>
```

## Response Body

```xml
<?xml version="1.0" encoding="utf-8"?>
<aunit:runResult xmlns:aunit="http://www.sap.com/adt/aunit">
    <external>
        <coverage adtcore:uri="/sap/bc/adt/runtime/traces/coverage/measurements/0800279158BE1ED8AFFA4794B38F5A3A" xmlns:adtcore="http://www.sap.com/adt/core"/>
    </external>
    <alerts/>
    <program adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext" adtcore:type="CLAS/OC" adtcore:name="ZCL_ZSADL_TEST_DPC_EXT" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core">
        <alerts/>
        <testClasses>
            <testClass adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/includes/testclasses#start=12,6" adtcore:type="CLAS/OCN/testclasses" adtcore:name="LTCL_TEST" adtcore:packageName="$SADL_TEST">
                <alerts/>
                <testMethods>
                    <testMethod adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/includes/testclasses#start=14,9" adtcore:type="CLAS/OCN/testclasses" adtcore:name="DUMMY" adtcore:packageName="$SADL_TEST" executionTime="0" unit="s">
                        <alerts/>
                    </testMethod>
                </testMethods>
            </testClass>
        </testClasses>
    </program>
</aunit:runResult>
```

# Code Coverage

## Request Body

```xml
<?xml version="1.0" encoding="UTF-8"?><cov:query xmlns:cov="http://www.sap.com/adt/cov" xmlns:adtcore="http://www.sap.com/adt/core">
  <adtcore:objectSets xmlns:adtcore="http://www.sap.com/adt/core">
    <objectSet kind="inclusive">
      <adtcore:objectReferences>
        <adtcore:objectReference adtcore:uri="/sap/bc/adt/vit/wb/object_type/devck/object_name/$SADL_TEST"/>
      </adtcore:objectReferences>
    </objectSet>
  </adtcore:objectSets>
</cov:query>
```

## Response Body

```xml
<?xml version="1.0" encoding="utf-8"?>
<cov:result name="ADT_ROOT_NODE" xmlns:cov="http://www.sap.com/adt/cov">
    <atom:link href="/sap/bc/adt/runtime/traces/coverage/results/0800279158BE1ED8AFFA48A18279FA3A" rel="self" xmlns:atom="http://www.w3.org/2005/Atom"/>
    <atom:link href="/sap/bc/adt/runtime/traces/coverage/results/0800279158BE1ED8AFFA48A18279FA3A/statements" rel="http://www.sap.com/adt/relations/runtime/traces/coverage/results/bulkstatements" xmlns:atom="http://www.w3.org/2005/Atom"/>
    <nodes>
        <node>
            <adtcore:objectReference adtcore:uri="/sap/bc/adt/vit/wb/object_type/devck/object_name/%24SADL_TEST" adtcore:type="DEVC/K" adtcore:name="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
            <coverages>
                <coverage type="branch" total="95" executed="0"/>
                <coverage type="procedure" total="43" executed="0"/>
                <coverage type="statement" total="1302" executed="0"/>
            </coverages>
            <nodes>
                <node>
                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=266,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_DPC============CP" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                    <coverages>
                        <coverage type="branch" total="85" executed="0"/>
                        <coverage type="procedure" total="35" executed="0"/>
                        <coverage type="statement" total="154" executed="0"/>
                    </coverages>
                    <nodes>
                        <node>
                            <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=266,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_DPC" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                            <coverages>
                                <coverage type="branch" total="85" executed="0"/>
                                <coverage type="procedure" total="35" executed="0"/>
                                <coverage type="statement" total="154" executed="0"/>
                            </coverages>
                            <atom:link href="/sap/bc/adt/runtime/traces/coverage/results/0800279158BE1ED8AFFA48A18279FA3A/statements/ADT_ROOT_NODE/%24SADL_TEST/ZCL_ZSADL_TEST_DPC%3d%3d%3d%3d%3d%3d%3d%3d%3d%3d%3d%3dCP/ZCL_ZSADL_TEST_DPC" rel="http://www.sap.com/adt/relations/runtime/traces/coverage/results/statements" type="application/xml+scov" xmlns:atom="http://www.w3.org/2005/Atom"/>
                            <nodes>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=269,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~CREATE_DEEP_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=278,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~CREATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="7" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="13" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=381,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~DELETE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="7" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="10" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=451,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~EXECUTE_ACTION" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=457,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~GET_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="13" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="19" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=576,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~GET_ENTITYSET" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="7" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="13" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=703,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~GET_EXPANDED_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=718,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~GET_EXPANDED_ENTITYSET" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=738,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~PATCH_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="3" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="6" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=761,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_MGW_APPL_SRV_RUNTIME~UPDATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="13" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="19" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=877,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_SB_DPC_COMM_SERVICES~COMMIT_WORK" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="5" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="8" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=916,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_SB_DPC_COMM_SERVICES~GET_GENERATION_STRATEGY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=922,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_SB_DPC_COMM_SERVICES~LOG_MESSAGE" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="4" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=945,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_SB_DPC_COMM_SERVICES~RFC_EXCEPTION_HANDLING" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="3" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=959,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_SB_DPC_COMM_SERVICES~RFC_SAVE_LOG" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="4" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=978,9" adtcore:type="CLAS/OM/public" adtcore:name="/IWBEP/IF_SB_DPC_COMM_SERVICES~SET_INJECTION" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="3" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="4" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=988,9" adtcore:type="CLAS/OM" adtcore:name="CHECK_SUBSCRIPTION_AUTHORITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=996,9" adtcore:type="CLAS/OM/public" adtcore:name="IF_SADL_GW_DPC_UTIL~GET_DPC" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="5" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1131,9" adtcore:type="CLAS/OM/public" adtcore:name="IF_SADL_GW_EXTENSION_CONTROL~SET_EXTENSION_MAPPING" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1137,9" adtcore:type="CLAS/OM/public" adtcore:name="IF_SADL_GW_QUERY_CONTROL~SET_QUERY_OPTIONS" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1143,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005SET_CREATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1150,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005SET_DELETE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1155,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005SET_GET_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1161,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005SET_GET_ENTITYSET" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1168,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005SET_UPDATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1175,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005TSET_CREATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1182,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005TSET_DELETE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1187,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005TSET_GET_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1193,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005TSET_GET_ENTITYSET" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1200,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005TSET_UPDATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1207,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005_WITH_TEXTSE_CREATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1214,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005_WITH_TEXTSE_DELETE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1219,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005_WITH_TEXTSE_GET_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1225,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005_WITH_TEXTSE_GET_ENTITYSET" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc/source/main#start=1232,9" adtcore:type="CLAS/OM/protected" adtcore:name="T005_WITH_TEXTSE_UPDATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="2" executed="0"/>
                                    </coverages>
                                </node>
                            </nodes>
                        </node>
                    </nodes>
                </node>
                <node>
                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/source/main#start=14,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_DPC_EXT========CP" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                    <coverages>
                        <coverage type="branch" total="1" executed="0"/>
                        <coverage type="procedure" total="1" executed="0"/>
                        <coverage type="statement" total="3" executed="0"/>
                    </coverages>
                    <nodes>
                        <node>
                            <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/source/main#start=14,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_DPC_EXT" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                            <coverages>
                                <coverage type="branch" total="1" executed="0"/>
                                <coverage type="procedure" total="1" executed="0"/>
                                <coverage type="statement" total="3" executed="0"/>
                            </coverages>
                            <atom:link href="/sap/bc/adt/runtime/traces/coverage/results/0800279158BE1ED8AFFA48A18279FA3A/statements/ADT_ROOT_NODE/%24SADL_TEST/ZCL_ZSADL_TEST_DPC_EXT%3d%3d%3d%3d%3d%3d%3d%3dCP/ZCL_ZSADL_TEST_DPC_EXT" rel="http://www.sap.com/adt/relations/runtime/traces/coverage/results/statements" type="application/xml+scov" xmlns:atom="http://www.w3.org/2005/Atom"/>
                            <nodes>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/source/main#start=17,9" adtcore:type="CLAS/OM" adtcore:name="T005TSET_UPDATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="3" executed="0"/>
                                    </coverages>
                                </node>
                            </nodes>
                        </node>
                    </nodes>
                </node>
                <node>
                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=67,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_MPC============CP" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                    <coverages>
                        <coverage type="branch" total="9" executed="0"/>
                        <coverage type="procedure" total="7" executed="0"/>
                        <coverage type="statement" total="1145" executed="0"/>
                    </coverages>
                    <nodes>
                        <node>
                            <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=67,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_MPC" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                            <coverages>
                                <coverage type="branch" total="9" executed="0"/>
                                <coverage type="procedure" total="7" executed="0"/>
                                <coverage type="statement" total="1145" executed="0"/>
                            </coverages>
                            <atom:link href="/sap/bc/adt/runtime/traces/coverage/results/0800279158BE1ED8AFFA48A18279FA3A/statements/ADT_ROOT_NODE/%24SADL_TEST/ZCL_ZSADL_TEST_MPC%3d%3d%3d%3d%3d%3d%3d%3d%3d%3d%3d%3dCP/ZCL_ZSADL_TEST_MPC" rel="http://www.sap.com/adt/relations/runtime/traces/coverage/results/statements" type="application/xml+scov" xmlns:atom="http://www.w3.org/2005/Atom"/>
                            <nodes>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=70,9" adtcore:type="CLAS/OM" adtcore:name="DEFINE" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="6" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=88,9" adtcore:type="CLAS/OM/private" adtcore:name="DEFINE_ASSOCIATIONS" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="7" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=140,9" adtcore:type="CLAS/OM/private" adtcore:name="DEFINE_T005" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="405" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=708,9" adtcore:type="CLAS/OM/private" adtcore:name="DEFINE_T005T" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="69" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=832,9" adtcore:type="CLAS/OM/private" adtcore:name="DEFINE_T005_WITH_TEXT" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="470" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=1486,9" adtcore:type="CLAS/OM" adtcore:name="GET_LAST_MODIFIED" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="3" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="4" executed="0"/>
                                    </coverages>
                                </node>
                                <node>
                                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc/source/main#start=1504,9" adtcore:type="CLAS/OM/public" adtcore:name="LOAD_TEXT_ELEMENTS" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                                    <coverages>
                                        <coverage type="branch" total="1" executed="0"/>
                                        <coverage type="procedure" total="1" executed="0"/>
                                        <coverage type="statement" total="184" executed="0"/>
                                    </coverages>
                                </node>
                            </nodes>
                        </node>
                    </nodes>
                </node>
                <node>
                    <adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc_ext/source/main#start=13,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_MPC_EXT========CP" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_mpc_ext" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/>
                    <coverages>
                        <coverage type="branch" total="0" executed="0"/>
                        <coverage type="procedure" total="0" executed="0"/>
                        <coverage type="statement" total="0" executed="0"/>
                    </coverages>
                </node>
            </nodes>
        </node>
    </nodes>
</cov:result>
```
