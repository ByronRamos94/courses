package com.observer;

public class CurrentConditionsDisplay implements Observer, DisplayElement{
    private float temperature;
    private float humidity;
    private WeatherData weatherData;

    public CurrentConditionsDisplay(WeatherData weatherData) {
        weatherData = weatherData;
        weatherData.registerObserver(this);
    }

    @Override
    public void update(float temperature, float humidity, float pressure) {
    this.humidity = humidity;
    this.temperature = temperature;
    display();
    }

    @Override
    public void display() {
        System.out.println("Current Conditions: " + temperature + "F" + "and " + humidity + " humidity");
    }
}
