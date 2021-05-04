__SINGLETON PATTERN__

**MOTIVATION**

- Some items make sense to only have 1 in the system
1. Database Repository
2. Object Factory

- This may be done because the construction of the component is very expensive
- We only do this once
- Give everyone the same instance
- Shouldnt have to create copies of the object
- Need to take care of the lazy instantiation

