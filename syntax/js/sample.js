/*
 * The purpose of the file is to use all available syntax tokens in JS
 */

// import
import theDefault, { named1, named2 } from 'src/mylib';
import theDefault from 'src/mylib';
import { named1, named2 } from 'src/mylib';
import { named1 as myNamed1, named2 } from 'src/mylib';
import * as mylib from 'src/mylib';
import 'src/mylib';

export default foo;

// declarations and assignations
g = 'global';
var v = 'v';
let l = 'l';
const C = 'C';

// primitives
var undef = undefined;
var nil = null;
var bool = true;
var number = 42;
var nan = NaN;
var sString = 'string';
var dString = "string";

// complex
var array = [0, 4, "foo"];
var object = {
    "foo": "bar",
    qux: 42
};

// function declaration with default params, rest params
function foo (bar, buz = 'def', ...rest) {

    // conditions
    if (1 == 2 && 3 <= 4) {
        return 'wot?';
    } else if {
        console.log(arguments);
    }

    switch (bar) {
        case '1':
            powipu = void 0;
        break;

        default:
        break;
    }

    // loops
    for (let i = 0; i <= 42; i++) {
        // arithmetic
        bar = bar + buz + 2;

        // interpolation
        qux = `no ${i} sub`;
        continue;
    }

    do {
        "hop".slice(2);
    } while (1 > 3)

    // exceptions
    try {
        throw new Exception('error');
    } catch(e) {
        console.error(e);
    } finally {

    }

    // function expression
    var qux = function yep (hello) {
        this.inside = hello;
        return typeof '1' === 'string';
    };

    // fat arrow
    var fiz = x => x + x;

    // generator
    function* generator() {
       yield 'res';
    }
}

foo.prototype = {};

// class
class Foo extends Bar {

    constructor() {
        super();
    }

    small(props) {
        // destructuration
        var { a } = props;
    }
}
