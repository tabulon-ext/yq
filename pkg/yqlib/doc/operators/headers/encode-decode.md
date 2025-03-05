# Encoder / Decoder

Encode operators will take the piped in object structure and encode it as a string in the desired format. The decode operators do the opposite, they take a formatted string and decode it into the relevant object structure.

Note that you can optionally pass an indent value to the encode functions (see below).

These operators are useful to process yaml documents that have stringified embedded yaml/json/props in them.


| Format | Decode (from string) | Encode (to string) |
| --- | -- | --|
| Yaml | from_yaml/@yamld | to_yaml(i)/@yaml |
| JSON | from_json/@jsond | to_json(i)/@json |
| Properties | from_props/@propsd  | to_props/@props |
| CSV | from_csv/@csvd | to_csv/@csv |
| TSV | from_tsv/@tsvd | to_tsv/@tsv |
| XML | from_xml/@xmld | to_xml(i)/@xml |
| Base64 | @base64d | @base64 |
| URI | @urid | @uri |
| Shell |  | @sh |


See CSV and TSV [documentation](https://mikefarah.gitbook.io/yq/usage/csv-tsv) for accepted formats.

XML uses the `--xml-attribute-prefix` and `xml-content-name` flags to identify attributes and content fields.


Base64 assumes [rfc4648](https://rfc-editor.org/rfc/rfc4648.html) encoding. Encoding and decoding both assume that the content is a utf-8 string and not binary content.
