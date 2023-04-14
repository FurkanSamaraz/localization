# Core - Utilities

Provides common utility functions for a wide range of purposes. Some of the included functions are:

### Fuzzy Search

Providing `FuzzyDistance()` and `FuzzyDistanceName()` for calculating the Levenshtein distance between two strings. The "Name" variant does this by taking into account western Names of family name last and reversing the order to get slightly better results for people names.

### Language

Provides `ParseLanguage()` and `ParseLanguageHTTPFallback()` functions for attempting to find a BSP-47 language tag that the system currently supports. At this time, only English and Turkish are provided hardcoded.

### Parsing

Helpful function `ParseBool()` takes and input string and returns a boolean true if it starts with "t" or "1". Used for attempting to interpret HTTP Query parameters. Checks the first letter, so even if the word "Turnip" is provided, it will equate to true. Also note, an empty string returns false.

### Password

Currently `ValidatePassword()` just checks the length of the password to be above 5. In the future a list of dumb passwords will be checked as well to increase security. We are not a fan of stupidly complicated password patterns, they offer little increase to password strength and just makes them forgettable. 

### Random

Using `GenerateRandomCode()` a randomly selected Base32 string of given length is returned.

### Slices

Helpful function `InterfaceSlice()` takes an interface object and uses reflection to return an actual slice of interface objects. Go is quirky like this sometimes.

A fairly complicated function `FindSliceDifferences()` will compare two given slices using a provided equates function, and will return 3 slices of additions, subtractions, and same objects. Additions are from the perspective that the new slice (second parameter)  has an object that the old slice (first parameter) didn't. The inverse applies for the subtractions. The equates function has the signature `func(old, new int)bool` and is provided with the slice indices for objects being compared. Should return true if the objects equate.

Instead of splitting up the slices, `ForSliceDifferences()` will perform the same comparison, but will execute provided call-back functions when a difference is encountered.

If comparison isn't needed, the `FindInSlice()` function will search a slice for an entry. Again, it uses an equate function, but will return the object when found.

### Strings

The function `TitleCaseString()` takes an incoming string, and performs English title-casing for proper nouns. Used for cleaning up user inputs for titles. Essentially just capitalizes the first letter in each word.

### Time

Because the built-in `time` package doesn't provide helpful date facilities some are provided. The `DateEqual()` function takes two `time.Time` objects and equates only the date portion of them.

`TimeToDate()` rounds down a `time.Time` object by setting the time components to 0.

`DaysBetweenDates()` takes two `time.Time` objects and computes the number of whole days between them.