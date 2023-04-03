import { useState } from "react";

import prettyBytes from "pretty-bytes";

import Layout from "@/components/Layout";
import FileUploader from "@/components/FileUploader";
import { useSettingsContext } from "@/context/SettingsProvider";
import siteConfig from "site.config";

export default function Home() {
	const [currentFile, setCurrentFile] = useState<any>();

	const { settings } = useSettingsContext();

	const handleResult = (convertedImage: ArrayBuffer, previousImage: File) => {
		const filename = previousImage.name.replace(
			/\.[^/.]+$/,
			`.${settings.fileOutputId}`
		);
		const newImage = new File([convertedImage], filename, {
			type: `image/${settings.fileOutputId}`,
		});
		const imageURL = URL.createObjectURL(newImage);

		setCurrentFile({
			name: filename,
			size: prettyBytes(newImage.size),
			source: imageURL,
			information: {
				Created: new Date(newImage.lastModified).toDateString(),
				"Last modified": new Date(newImage.lastModified).toDateString(),
			},
		});
	};

	return (
		<Layout>
			<div className="flex flex-1 items-stretch">
				<main className="flex-1 overflow-y-auto">
					<div className="mx-auto max-w-7xl px-4 pt-8 sm:px-6 lg:px-8">
						<div className="border-b border-gray-200 pb-5">
							<h1 className="text-base text-xl font-semibold leading-6 text-gray-900">
								{siteConfig.slogan}
							</h1>
							<p className="mt-2 max-w-4xl text-sm text-gray-500">
								Convert your HEIC photos to JPEGs in seconds with SwiftConvert -
								the fast and free online tool.
							</p>
							<p className="mt-4 max-w-4xl text-sm text-gray-500">
								Currently supported output formats:{" "}
								<span className="text-blue-800">.png | .jpeg</span>
							</p>
						</div>
						<section className="mt-2 pb-16" aria-labelledby="main-heading">
							<h2 id="main-heading" className="sr-only">
								Converting photos
							</h2>
							<FileUploader handleResult={handleResult} />
						</section>
					</div>
				</main>
				{/* Details sidebar */}
				{currentFile && (
					<aside className="w-96 h-screen overflow-y-auto border-l border-gray-200 bg-white p-8 lg:block">
						<div className="space-y-6 pb-12 h-96">
							<div>
								<div className="aspect-h-7 aspect-w-10 block w-full overflow-hidden rounded-lg">
									<img
										src={currentFile.source}
										alt={currentFile.name}
										className="object-cover"
									/>
								</div>
								<div className="mt-4 flex items-start justify-between">
									<div>
										<h2 className="text-lg font-medium text-gray-900">
											<span className="sr-only">Details for </span>
											{currentFile.name}
										</h2>
										<p className="text-sm font-medium text-gray-500">
											{currentFile.size}
										</p>
									</div>
								</div>
							</div>
							<div>
								<h3 className="font-medium text-gray-900">Information</h3>
								<dl className="mt-2 divide-y divide-gray-200 border-b border-t border-gray-200">
									{Object.keys(currentFile.information).map((key) => (
										<div
											key={key}
											className="flex justify-between py-3 text-sm font-medium"
										>
											<dt className="text-gray-500">{key}</dt>
											<dd className="whitespace-nowrap text-gray-900">
												{currentFile.information[key]}
											</dd>
										</div>
									))}
								</dl>
							</div>
							<div className="flex gap-x-3">
								<a
									type="button"
									download={currentFile.name}
									href={currentFile.source}
									title={currentFile.name}
									className="flex-1 rounded-md bg-blue-600 px-3 py-2 text-sm text-center font-semibold text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600"
								>
									Download
								</a>
								<button
									type="button"
									onClick={() => setCurrentFile(undefined)}
									className="flex-1 rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"
								>
									Delete
								</button>
							</div>
						</div>
					</aside>
				)}
			</div>
		</Layout>
	);
}
