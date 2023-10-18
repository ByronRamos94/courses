package com.testing;

public class Scar implements Arma{
    @Override
    public void atacar() {
        System.out.println("disparando como una scar");
    }

    @Override
    public void recargar() {
        System.out.println("recargando 8 balas");

    }
}
