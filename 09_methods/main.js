class Fruit {
  constructor(name, count) {
    this.name = name;
    this.count = count;
  }

  print() {
    console.log({
      name: this.name,
      count: this.count
    });
  }
}

let a = new Fruit("Apple", 12);
a.print();
