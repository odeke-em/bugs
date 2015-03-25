util = require('util');

// An excerpt from a package
function testf(indexName, typeName, documentId, options, cb) {
    console.log('originally\n\t', typeof arguments, util.inspect(arguments));

    cb = arguments[arguments.length - 1];
    console.log('cb before replacement\n\t', util.inspect(cb));

    arguments[arguments.length - 1] = false;
    console.log('after replacement\n\t', cb, arguments);
};

console.log('#');
testf('flux', 'gox', 'dome', {}, {2: 10});
console.log('--');
testf();
console.log('#');

/*
 * node --version 0.10.32
#
originally
	 object { '0': 'flux', '1': 'gox', '2': 'dome', '3': {}, '4': { '2': 10 } }
cb before replacement
	 { '2': 10 }
after replacement
	 false { '0': 'flux', '1': 'gox', '2': 'dome', '3': {}, '4': false }
--
originally
	 object {}
cb before replacement
	 undefined
after replacement
	 undefined { '-1': false }
#
*/

/*
 * node --version 0.13.0
#
originally
	 object { '0': 'flux', '1': 'gox', '2': 'dome', '3': {}, '4': { '2': 10 } }
cb before replacement
	 { '2': 10 }
after replacement
	 false { '0': 'flux', '1': 'gox', '2': 'dome', '3': {}, '4': false }
--
originally
	 object {}
cb before replacement
	 undefined
after replacement
	 undefined { '-1': false }
#
*/

/*
 * node --version 0.8.20
#
originally
	 object { '0': 'flux', '1': 'gox', '2': 'dome', '3': {}, '4': { '2': 10 } }
cb before replacement
	 { '2': 10 }
after replacement
	 false { '0': 'flux', '1': 'gox', '2': 'dome', '3': {}, '4': false }
--
originally
	 object {}
cb before replacement
	 undefined
after replacement
	 undefined { '-1': false }
#
*/
