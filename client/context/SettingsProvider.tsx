import React, { createContext, useContext, useState, useEffect } from "react";

import { FileType, SettingsKeys } from "@/types/index";

type Settings = { [key: string]: string };

type SettingsContextProperties = {
	settings: Settings | null;
	updateSettings: (fileInputId: string, fileOutputId: string) => void;
};

const SettingsContext = createContext<SettingsContextProperties>({
	settings: null,
	updateSettings: (fileInputId, fileOutputId) => undefined,
});

interface Properties {
	children: React.ReactNode;
}

const SettingsProvider = ({ ...properties }: Properties) => {
	const [settings, setSettings] = useState<Settings>(null);

	const updateSettings = (fileInputId: string, fileOutputId: string) => {
		localStorage.setItem(SettingsKeys.FILE_INPUT_TYPE, fileInputId);
		localStorage.setItem(SettingsKeys.FILE_OUTPUT_TYPE, fileOutputId);

		setSettings({
			fileInputId: fileInputId,
			fileOutputId: fileOutputId,
		});
	};

	const returnValue = {
		settings,
		updateSettings,
	};

	useEffect(() => {
		if (!settings) {
			const input = localStorage.getItem(SettingsKeys.FILE_INPUT_TYPE);
			const output = localStorage.getItem(SettingsKeys.FILE_OUTPUT_TYPE);

			if (!input && output) {
				updateSettings(FileType.heic, output);
			} else if (input && !output) {
				updateSettings(input, FileType.jpeg);
			} else {
				updateSettings(FileType.heic, FileType.jpeg);
			}
		}
	}, []);

	return <SettingsContext.Provider value={returnValue} {...properties} />;
};

const useSettingsContext = () => {
	const context = useContext(SettingsContext);
	if (!context)
		throw new Error(
			"useSettingsContext must be used within a SettingsProvider"
		);

	return context;
};

export {
	SettingsContext,
	SettingsProvider as SettingsProvider,
	useSettingsContext,
};
