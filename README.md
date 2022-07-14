# Information Retrieval - basic algorithms

This project inludes three basic algorithms:
- Boolean retrieval
- Skip list (intersect with skip pointers)
- Positional intersect

## How it works?

First of all, you need to specfiy a path where your documents are kept. You can do this by using this the below flag:
```
ir-project -path="~/documents/"
```

Then you need to choose one of the three algorithms by these flags: _boolean_, _skip_ and _positional_.

Now let's see how each one works:

### boolean
```
ir-project -path="~/documents/" -boolean "term1" "term2" "term3"
```

Please notice that it only supports `AND` query.

### Skip
```
ir-project -path="~/documents/" -skip "term1" "term2" "term3"
```

Please notice that it only supports `AND` query.

### positional
```
ir-project -path="~/documents/" -positional -phrase="Information retrieval is exciting"
```

With positional intersect algorithm, our aim is to be able to search for a whole phrase instead of checking existense of some words in documents. That's why it gives a phrase as an input.