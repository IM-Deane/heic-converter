export interface Input {
	id: string | number | FileType;
	name: string;
	unavailable?: boolean;
}

export enum SettingsKeys {
	FILE_INPUT_TYPE = "fileInput",
	FILE_OUTPUT_TYPE = "fileOutput",
}

export enum FileType {
	heic = "heic",
	jpeg = "jpeg",
	png = "png",
}

export const fileTypes: Input[] = [
	{ id: "jpeg", name: ".JPEG", unavailable: false },
	{ id: "png", name: ".PNG", unavailable: false },
];

export const settingsInputTypes: Input[] = [
	{ id: "heic", name: ".HEIC", unavailable: false },
	{ id: "jpeg", name: ".JPEG", unavailable: true },
	{ id: "png", name: ".PNG", unavailable: true },
];
