// For loops.
for (let i = 0; i < 10; i++) {
  console.log(i);
}

// For loop but with while.
let i = 0;
while (i < 10) {
  console.log(i);
  i++;
}

// Infinite loops with while.
i = 0;
while (true) {
  console.log(i++);
  if (i > 10) {
    break;
  }
}

// Do while loops from the grave.
i = 0;
do {
  console.log(i);
  i++;
} while (i < 10);

// Traversing arrays and objects.

// let p = [0, 1, 2];
let p = {
  a: 1,
  b: 2,
  c: 3
};

// key, value loop
for (var key in p) {
  if (p.hasOwnProperty(key)) {
    console.log(key + " -> " + p[key]);
  }
}
