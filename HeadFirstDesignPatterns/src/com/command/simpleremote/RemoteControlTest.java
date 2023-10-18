package com.command.simpleremote;

import com.command.simpleremote.commands.GarageDoorOpenCommand;
import com.command.simpleremote.commands.LightOnCommand;
import com.command.simpleremote.devices.GarageDoor;
import com.command.simpleremote.devices.Light;

public class RemoteControlTest {
	public static void main(String[] args) {
		SimpleRemoteControl remote = new SimpleRemoteControl();
		Light light = new Light();
		GarageDoor garageDoor = new GarageDoor();
		LightOnCommand lightOn = new LightOnCommand(light);
		GarageDoorOpenCommand garageOpen =
		    new GarageDoorOpenCommand(garageDoor);
 
		remote.setCommand(lightOn);
		remote.buttonWasPressed();
		remote.setCommand(garageOpen);
		remote.buttonWasPressed();
    }
	
}
