__BUILDER DESIGN PATTERN__
- Some objects are small and can be called in a single constructor call
- Some require a lot of ceremony to make
- This enforces a piecewise approach at constructing complex objects
- Provides an API for constructing an object step-by-step

We can use a fluent interface in GO to chain calls together.
We return the item that is being modified in the function
so that we can chain subsequent calls
