package com.observer;

import java.util.ArrayList;
import java.util.List;

public class WeatherData implements Subject{
    private float humidity;
    private float temperature;
    private float pressure;
    private List<Observer> observers;


    public WeatherData() {
        observers = new ArrayList<Observer>();
    }

    @Override
    public void registerObserver(Observer o) {
        observers.add(o);
    }

    @Override
    public void removeObserver(Observer o) {
        observers.remove(o);
    }

    @Override
    public void notifyObservers() {
        for (Observer observer : observers){
            observer.update(temperature, humidity, pressure);
        }
    }

    public void measurementsChanged(){
        notifyObservers();

    };
    public void setMeasurements(float humidity, float temperature, float pressure ){
        this.humidity = humidity;
        this.temperature = temperature;
        this.pressure = pressure;
        measurementsChanged();

    }
    public float getHumidity() {
        return humidity;
    }

    public float getTemperature() {
        return temperature;
    }

    public float getPressure() {
        return pressure;
    }
}
