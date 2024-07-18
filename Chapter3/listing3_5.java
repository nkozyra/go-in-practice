public class MyApp {
    public static void main(String[] args) {
        Animal a = new Animal();
        a.speak();

        Cat c = new Cat();
        c.speak();

        Dog d = new Dog();
        d.speak();

        Llama l = new Llama();
        l.speak();

    }
}

class Animal {
    void speak() {
        System.out.println("nondescript animal noise?");
    }
}

class Cat extends Animal {
    void speak() {
        System.out.println("meow");
    }
}

class Dog extends Animal {
    void speak() {
        System.out.println("woof");
    }
}

class Llama extends Animal { }
