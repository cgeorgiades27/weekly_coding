def flatten_dict(nested):
    dict_iter = list(nested.items())
    flat_dict = {}
    while dict_iter:
        key, val = dict_iter.pop()
        if isinstance(val, dict):
            for nested_key, nested_val in val.items():
                dict_iter.append((f"{key}.{nested_key}", nested_val))
        else:
            flat_dict[key] = val
    return flat_dict
    

d = {
    "key": 3,
    "foo": {
        "a": 5,
        "bar": {
            "baz": 8
        }
    }
}   

print(flatten_dict(d))
