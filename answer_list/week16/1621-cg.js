/*
{
    "key": 3,
    "foo": {
        "a": 5,
        "bar": {
            "baz": 8
        }
    }
}
```
it should become:
```json
{
    "key": 3,
    "foo.a": 5,
    "foo.bar.baz": 8
}
*/

let expanded = '{ "key": 3, "foo": { "a": 5, "bar": { "baz": 8 } } }';

function flatten(obj) {
  if (typeof obj != "object") return;
  for (i in obj) {
    flatten(obj[i]);
  }
}

flatten(JSON.parse(expanded));
