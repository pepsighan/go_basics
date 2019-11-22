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
