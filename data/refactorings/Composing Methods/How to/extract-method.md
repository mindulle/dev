# Benefits

By minimizing the number of unneeded methods, you make the code more straightforward.

# How to Refactor

1. Make sure that the method isn’t redefined in subclasses. If the method is redefined, refrain from this technique.

2. Find all calls to the method. Replace these calls with the content of the method.

3. Delete the method.
