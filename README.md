# Shape sorter code using go

Takes input of shapes and output arrays of shapes sorted based on
1. A shape is of type circle or rectangle
2. It is identified by its area (computed) and color (provided as input). Circle requires radius and color whereas rectangle requires length, breadth and color.
3. Color can be of type Red, Green, Blue only.
4. Input given is mixture of circles and rectangles, the program outputs 2 distinct array containing sorted circles and rectangles.
5. The sorting order should be increasing area followed by configurable color order.

## Methods to run the program

1. `CLI` - accepts raw json string as input. For malformed input, err on stderr and non zero exit status code.
2. `FILE` - accepts filename and uses that file contents as json input. For malformed input/ file not found, err on stderr and non zero exit status code. 
3. `WEB` - web service which accepts a port number. The /shape-sort route is utilized which accepts json post input and provides response in return. For malformed input, response header is 400, body contains failure reason.

## Example of input 
A sample Input json may look like this:

{
    "rectangles": [
        { "length": 10.0, "breadth": 20.0, color: “red”},
        { "length": 20.0, "breadth": 30.0, color: “blue”},
        { "length": 20.0, "breadth": 30.0, color: “green”},
        { "length": 5.0, "breadth": 15.0, color: “green”}  
    ],
    "circles": [
        { "radius": 20.0, color: “red” },
        { "radius": 15.0, color: “blue” },
        { "radius": 30.0, color: “red” }     
    ]
}

## Example of output
CLI:

shape_shifter --json='{}'

output

-----------------------

Rectangles:

{ "length": 5.0, "breadth": 15.0, color: “green”}

{ "length": 10.0, "breadth": 20.0, color: “red”}

{ "length": 20.0, "breadth": 30.0, color: “green”}

{ "length": 20.0, "breadth": 30.0, color: “blue”}



Circles:

{ "radius": 15.0, color: “blue” }

{ "radius": 20.0, color: “red” }

{ "radius": 30.0, color: “red” }



FILE:

shape_shifter --json-file=/path/to.json

output

-----------------------

Rectangles:

{ "length": 5.0, "breadth": 15.0, color: “green”}

{ "length": 10.0, "breadth": 20.0, color: “red”}

{ "length": 20.0, "breadth": 30.0, color: “green”}

{ "length": 20.0, "breadth": 30.0, color: “blue”}



Circles:

{ "radius": 15.0, color: “blue” }

{ "radius": 20.0, color: “red” }

{ "radius": 30.0, color: “red” }



WEB:

shape_shifter --port=8081

curl -XPOST ‘http://localhost:8081/shape-sort’ -d '{}'

output

-----------------------

{

“rectangles: [

{ "length": 5.0, "breadth": 15.0, color: “green”},

{ "length": 10.0, "breadth": 20.0, color: “red”},

{ "length": 20.0, "breadth": 30.0, color: “green”},

{ "length": 20.0, "breadth": 30.0, color: “blue”}

],

“circles”: [

{ "radius": 15.0, color: “blue” },

{ "radius": 20.0, color: “red” },

{ "radius": 30.0, color: “red” }

]

}