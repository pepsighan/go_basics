function errorableFunction() {
  throw Error("bad things happen");
}

try {
  let value = errorableFunction();
  // Do something with the value.
} catch (error) {
  console.log("Error occurred: ", error.message);
}
