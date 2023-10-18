package com.command.simpleremote.commands;

import com.command.simpleremote.devices.Light;

public class LightOnCommand implements Command {
	Light light;
  
	public LightOnCommand(Light light) {
		this.light = light;
	}
 
	public void execute() {
		light.on();
	}
}
