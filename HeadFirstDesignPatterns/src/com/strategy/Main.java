package com.strategy;

import com.strategy.Duck;
import com.strategy.FlyRocketPowered;
import com.strategy.MallarDuck;
import com.strategy.ModelDuck;

public class Main {

    public static void main(String[] args) {
	// write your code here
        Duck duck = new MallarDuck();
        ModelDuck modelDuck = new ModelDuck();
        modelDuck.performFly();
        modelDuck.performQuack();
        modelDuck.setFlyBehavior(new FlyRocketPowered());
        modelDuck.performFly();

    }
}
