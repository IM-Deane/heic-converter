import { useState, useRef, useEffect, Fragment } from "react";

import { Dialog, Transition } from "@headlessui/react";

import SelectInput from "./SelectInput";
import { useSettingsContext } from "@/context/SettingsProvider";
import { Input, fileTypes, settingsInputTypes } from "@/types/index";

export default function SettingsModal({ isOpen, setIsOpen }) {
	const { settings, updateSettings } = useSettingsContext();

	const [isReady, setIsReady] = useState(false);
	const [selectedInputType, setSelectedInputType] = useState<Input>(null);
	const [selectedOutputType, setSelectedOutputType] = useState<Input>(null);

	const cancelButtonRef = useRef(null);

	const handleSave = () => {
		updateSettings(
			selectedInputType.id as string,
			selectedOutputType.id as string
		);
		setIsOpen(false);
	};

	useEffect(() => {
		if (settings) {
			setSelectedInputType(
				settingsInputTypes.find((input) => input.id === settings?.fileInputId)
			);
			setSelectedOutputType(
				fileTypes.find((input) => input.id === settings?.fileOutputId)
			);
			setIsReady(true);
		}
	}, [settings]);

	return (
		<Transition.Root show={isOpen} as={Fragment}>
			<Dialog
				as="div"
				className="relative z-10 w-64"
				initialFocus={cancelButtonRef}
				onClose={() => setIsOpen(false)}
			>
				<Transition.Child
					as={Fragment}
					enter="ease-out duration-300"
					enterFrom="opacity-0"
					enterTo="opacity-100"
					leave="ease-in duration-200"
					leaveFrom="opacity-100"
					leaveTo="opacity-0"
				>
					<div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
				</Transition.Child>

				<div className="fixed inset-0 z-10 overflow-y-auto">
					<div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
						<Transition.Child
							as={Fragment}
							enter="ease-out duration-300"
							enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
							enterTo="opacity-100 translate-y-0 sm:scale-100"
							leave="ease-in duration-200"
							leaveFrom="opacity-100 translate-y-0 sm:scale-100"
							leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
						>
							<Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6">
								<form>
									<div className="space-y-12">
										<div className="grid grid-cols-1 gap-x-8 gap-y-10 border-b border-gray-900/10 pb-12 md:grid-cols-3">
											<div>
												<h2 className="text-base font-semibold leading-7 text-gray-900">
													Photo Settings
												</h2>
												<p className="mt-1 text-sm leading-6 text-gray-600">
													Modify the input and output file types
												</p>
											</div>

											{isReady && (
												<div className="grid max-w-2xl grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6 md:col-span-2">
													<div className="sm:col-span-3">
														<SelectInput
															inputLabel="Input"
															inputList={settingsInputTypes}
															selectedInput={selectedInputType}
															isDisabled
															handleSelectedInput={setSelectedInputType}
														/>
													</div>
													<div className="sm:col-span-3">
														<SelectInput
															inputLabel="Output"
															inputList={fileTypes}
															selectedInput={selectedOutputType}
															handleSelectedInput={setSelectedOutputType}
														/>
													</div>
												</div>
											)}
										</div>
									</div>
								</form>
								<div className="mt-5 sm:mt-6 sm:grid sm:grid-flow-row-dense sm:grid-cols-2 sm:gap-3">
									<button
										type="button"
										className="inline-flex w-full justify-center rounded-md bg-gray-800 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600 sm:col-start-2"
										onClick={handleSave}
									>
										Save
									</button>
									<button
										type="button"
										className="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:col-start-1 sm:mt-0"
										onClick={() => setIsOpen(false)}
										ref={cancelButtonRef}
									>
										Cancel
									</button>
								</div>
							</Dialog.Panel>
						</Transition.Child>
					</div>
				</div>
			</Dialog>
		</Transition.Root>
	);
}
