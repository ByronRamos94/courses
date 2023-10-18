package com.command.party.commands;

import com.command.party.commands.Command;
import com.command.party.devices.Light;

public class LivingroomLightOffCommand implements Command {
	Light light;

	public LivingroomLightOffCommand(Light light) {
		this.light = light;
	}
	public void execute() {
		light.off();
	}
	public void undo() {
		light.on();
	}
}
