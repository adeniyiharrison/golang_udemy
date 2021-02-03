# Ninja Level 5

1. Create your own type "person" which will have an underlying type of "struct" so that it can store the following data
    * First name
    * Last name
    * Favorite ice cream flavors

2. Take the code from the previous exercise, then store the values of type person in a map with key of last name. Access each value in the map. Print out the values, ranging over the slice.

3.
    * Create a new type: vehicle
        * The underlying type is a struct
        * The fields:
            * Doors
            * Color
    * Create two new types: Truck and Sedan
        * The underlying type is also struct
        * Embed the vehicle type
        * Give the truck struct a field call fourWheel and make it boolean
        * Give sedan struct luxury field also boolean
    * Using the vehicle, truck and sedan structs:
        * using a composite literal, create a value of truck and assign it values
        * Do same for sedan
    * Print out each of the fields
    * Print out a single field for these values

4. Create and use an anonymous struct