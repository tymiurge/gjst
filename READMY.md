# Objective
the idea is to have tool that generates json structure out of a template.
template itself would be a json file with specific syntax.

# Template Syntax
at first level, the json should contain <data> property as an embedded object, being actually the
template from which json is going to be generated.

all rest properties at the same level with the <data>, are going to be treated by the template engine as
variables; inside the <data> structure, the variables can be referred with {{@variableName}} syntax.

all properties prefixed with _ - like the <_1>, <_2> etc, etc, are going be ignored:
```json
{"_prefixedFieldName": "this is going to be treated kinda as a comment line"}
```

## referring syntax
referring syntax: @ and & symbols are used for variable references, $ - for referencing a built in function:
* @<variable name> refers to a variable defined in the variables scope, or in function scope; examples:
    -  @os: refers to a value of <os> field defined in the variables scope
        - @data_body(x): {key: {{ @x }}}
            - @data_body refers to a <data_body> field from variables scope; @x refers to the <x> variable defined as the <data_body>
    - &os: refers to a value <os> field that has been defined somewhere above the current field
    - $now(): refers to a built in function

## extracting syntax
inside the <data>, to define a property value it's possible to call a function with {{$funcName()}} syntax
preliminary list of supported functions:
  - now() - returns unix timestamp in milliseconds
  - timef(y, mm, d, h, mi, s, ms) - returns short date and time as a string of format 20130218T080910.123
  - epoch(of) - returns unix timestamp in milliseconds
    - of: string in format of short date and time up to ms, i.e: 20130218T080910.123
  - make(start, step, amount) - returns array of len equaled to <amount> where the first element is <start> and each subsequent one is a sum of its predecessor and the <step>
    - start {number}: first element of array
    - step {number}: value to be added to each element's predecessor to get the element value
    - amount {number}: amount of elements to be generated
    - example: $make(0, 1, 5) will generate [0, 1, 2, 3, 4]
  - key(of) - returns random key of the <of> json object
    - of {object}: json object
  - element(of) - returns random element of the <of> json array
    - of {array}: json array

## template structure  
  
```json
{
  "_025": "VARIABLES SCOPE",
  
  "os_def": {
    "ios": ["iOS 9", "iOS 10", "iOS 11", "iOS 12", "iOS 13"],
    "android": ["iOS 9", "iOS 10", "iOS 11", "iOS 12", "iOS 13"]
  },
  "gatheringInfoTimeStart": "20130218T080910.123",
  "gatheringInfoTimeEnd": "20130219T080910.123",
  
  
  "_026": "JSON TEMPLATE",
  
  "data": {
    "reportSendingTime": "{{ $now() }}",
    "reportGathering": "{{ $epoch($timef(2013, 2, 18, 12, 45, 30, 500) }}",
    "os": "{{ $key(@os_def) }}",
    "version": "{{ $element(@os_def[&os]) }}"
  }
}
```
  
## generating multiple json structures
for generating multiple json structures the data property value in template should be a loop:

```json
{
  "reportTimeArray": "{{ $make($epoch(20130218T080910.123, 10, 120) }}",
  "@data_body(x)": {
    "reportSendingTime": "{{ $now() }}",
    "reportGathering": "{{ @x }}"
  },
  "data": "{{ $for(@reportTimeArray, el => @data_body(@el) }}"
}
```
 