package com.strategy;

public class FlyRocketPowered implements FlyBehavior{

    @Override
    public void fly() {
        System.out.println("Am flying with a rocket");
    }
}
