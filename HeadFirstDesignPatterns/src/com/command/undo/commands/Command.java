package com.command.undo.commands;

public interface Command {
	public void execute();
	public void undo();
}
