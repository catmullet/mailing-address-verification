# Address Verification based on country and regex

### Add All country address regular expressions you want to the address_regex.xml file

```xml
<Countries>
    <Country>
        <Code>US</Code>
        <Name>United States</Name>
        <Regex>^\s*\S+(?:\s+\S+){2}</Regex>
        <MustValidate>true</MustValidate>
    </Country>
    <Country>
        <Code>DE</Code>
        <Name>Germany</Name>
        <Regex>^\s*\S+(?:\s+\S+){2}</Regex>
        <MustValidate>false</MustValidate>
    </Country>
</Countries>
```

### country = Either a country code or country name
### address = The street address to validate

### localhost:6001/address/verify/country/address
