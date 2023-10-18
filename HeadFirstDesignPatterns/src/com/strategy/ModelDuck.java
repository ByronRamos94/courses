package com.strategy;

public class ModelDuck extends Duck {
    public ModelDuck() {
        flyBehavior = new FlyNoWay();
        quackBehavior = new Squeack();
    }

    @Override
    public void display() {
        System.out.println("am a modelDuck");
    }
}
