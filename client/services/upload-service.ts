import axios, { AxiosResponse } from "axios";

import { FileType } from "@/types/index";

class UploadService {
	private service: any;

	constructor() {
		this.service = axios.create({
			baseURL: `${process.env.NEXT_PUBLIC_SSE_URL}/api`,
			headers: {
				"Content-Type": "multipart/form-data",
			},
		});
	}

	newImage = async (
		file: File,
		convertToFormat: FileType,
		onUploadProgress
	): Promise<AxiosResponse> => {
		const formData = new FormData();
		formData.append("file", file);

		return this.service.post("/convert", formData, {
			params: {
				format: convertToFormat,
			},
			responseType: "blob",
			onUploadProgress,
		});
	};
}

export default new UploadService();
