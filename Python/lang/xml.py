# coding=utf-8

import xml.etree.ElementTree as ET
import os
import sys

curDir = os.path.dirname(sys.argv[0])
xmlPath = os.path.join(curDir, "res", 'config.xml')
ET.register_namespace('', "http://www.w3.org/ns/widgets")
ET.register_namespace('cdv', "http://cordova.apache.org/ns/1.0")
tree = ET.parse(xmlPath)
root = tree.getroot()
# for child in root:
#     print(child.tag,":", child.attrib)
#     if 'name' in child.attrib and child.attrib['name'] == "Channel":
#         child.attrib['value'] = "10000"

preference = root.findall("{http://www.w3.org/ns/widgets}preference")
# print(preference)
for child in preference:
    print(child.tag, ":", child.attrib)
    if child.attrib['name'] == "Channel":
        child.attrib['value'] = "10000"
root.attrib["xmlns:cdv"] = "http://cordova.apache.org/ns/1.0"
tree.write(xmlPath, xml_declaration=True, encoding="utf-8", method='xml')
