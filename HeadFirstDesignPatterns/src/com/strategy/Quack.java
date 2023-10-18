package com.strategy;

public class Quack implements QuackBehavior{
    @Override
    public void quack() {
        System.out.printf("Quack");
    }
}
