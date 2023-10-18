package com.command.party.commands;

import com.command.party.commands.Command;
import com.command.party.devices.Stereo;

public class StereoOnWithCDCommand implements Command {
	Stereo stereo;
 
	public StereoOnWithCDCommand(Stereo stereo) {
		this.stereo = stereo;
	}
 
	public void execute() {
		stereo.on();
		stereo.setCD();
		stereo.setVolume(11);
	}

	public void undo() {
		stereo.off();
	}
}
