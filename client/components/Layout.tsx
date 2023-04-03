import { Fragment, useState } from "react";

import Link from "next/link";

import { Dialog, Menu, Transition } from "@headlessui/react";
import {
	Bars3BottomLeftIcon,
	CogIcon,
	PlusCircleIcon,
	HomeIcon,
	PhotoIcon,
	RectangleStackIcon,
	Squares2X2Icon as Squares2X2IconOutline,
	UserGroupIcon,
	XMarkIcon,
} from "@heroicons/react/24/outline";
import { MagnifyingGlassIcon, PlusIcon } from "@heroicons/react/20/solid";

import { classNames } from "@/utils/index";
import { useRouter } from "next/router";

const navigation = [
	// { name: "Home", href: "#", icon: HomeIcon, current: false },
	{ name: "Upload", href: "/upload", icon: PlusCircleIcon, current: false },
	{ name: "Photos", href: "/", icon: PhotoIcon, current: true },
	// { name: "Shared", href: "#", icon: UserGroupIcon, current: false },
	// { name: "Albums", href: "#", icon: RectangleStackIcon, current: false },
	// { name: "Settings", href: "#", icon: CogIcon, current: false },
];
const userNavigation = [
	{ name: "Your profile", href: "#" },
	// { name: "Sign out", href: "#" },
];

export default function Layout({ children }) {
	const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

	const router = useRouter();

	const handleSubmit = () => {
		console.log("Uploaded image");
	};

	navigation.forEach((item) => {
		if (item.href === router.pathname) {
			item.current = true;
		} else {
			item.current = false;
		}
	});

	return (
		<div className="flex">
			{/* Narrow sidebar */}
			<div className="hidden h-screen w-28 overflow-y-auto bg-indigo-700 md:block">
				<div className="flex w-full flex-col items-center py-6">
					<div className="flex flex-shrink-0 items-center">
						<img
							className="h-8 w-auto"
							src="https://tailwindui.com/img/logos/mark.svg?color=white"
							alt="Your Company"
						/>
					</div>
					<div className="mt-6 w-full flex-1 space-y-1 px-2">
						{navigation.map((item) => (
							<Link
								key={item.name}
								href={item.href}
								className={classNames(
									item.current
										? "bg-indigo-800 text-white"
										: "text-indigo-100 hover:bg-indigo-800 hover:text-white",
									"group flex w-full flex-col items-center rounded-md p-3 text-xs font-medium"
								)}
								aria-current={item.current ? "page" : undefined}
							>
								<item.icon
									className={classNames(
										item.current
											? "text-white"
											: "text-indigo-300 group-hover:text-white",
										"h-6 w-6"
									)}
									aria-hidden="true"
								/>
								<span className="mt-2">{item.name}</span>
							</Link>
						))}
					</div>
				</div>
			</div>

			{/* Mobile menu */}
			<Transition.Root show={mobileMenuOpen} as={Fragment}>
				<Dialog
					as="div"
					className="relative z-40 md:hidden"
					onClose={setMobileMenuOpen}
				>
					<Transition.Child
						as={Fragment}
						enter="transition-opacity ease-linear duration-300"
						enterFrom="opacity-0"
						enterTo="opacity-100"
						leave="transition-opacity ease-linear duration-300"
						leaveFrom="opacity-100"
						leaveTo="opacity-0"
					>
						<div className="fixed inset-0 bg-gray-600 bg-opacity-75" />
					</Transition.Child>

					<div className="fixed inset-0 z-40 flex">
						<Transition.Child
							as={Fragment}
							enter="transition ease-in-out duration-300 transform"
							enterFrom="-translate-x-full"
							enterTo="translate-x-0"
							leave="transition ease-in-out duration-300 transform"
							leaveFrom="translate-x-0"
							leaveTo="-translate-x-full"
						>
							<Dialog.Panel className="relative flex w-full max-w-xs flex-1 flex-col bg-indigo-700 pb-4 pt-5">
								<Transition.Child
									as={Fragment}
									enter="ease-in-out duration-300"
									enterFrom="opacity-0"
									enterTo="opacity-100"
									leave="ease-in-out duration-300"
									leaveFrom="opacity-100"
									leaveTo="opacity-0"
								>
									<div className="absolute right-0 top-1 -mr-14 p-1">
										<button
											type="button"
											className="flex h-12 w-12 items-center justify-center rounded-full focus:outline-none focus:ring-2 focus:ring-white"
											onClick={() => setMobileMenuOpen(false)}
										>
											<XMarkIcon
												className="h-6 w-6 text-white"
												aria-hidden="true"
											/>
											<span className="sr-only">Close sidebar</span>
										</button>
									</div>
								</Transition.Child>
								<div className="flex flex-shrink-0 items-center px-4">
									<img
										className="h-8 w-auto"
										src="https://tailwindui.com/img/logos/mark.svg?color=white"
										alt="Your Company"
									/>
								</div>
								<div className="mt-5 h-0 flex-1 overflow-y-auto px-2">
									<nav className="flex h-full flex-col">
										<div className="space-y-1">
											{navigation.map((item) => (
												<Link
													key={item.name}
													href={item.href}
													className={classNames(
														item.current
															? "bg-indigo-800 text-white"
															: "text-indigo-100 hover:bg-indigo-800 hover:text-white",
														"group flex items-center rounded-md py-2 px-3 text-sm font-medium"
													)}
													aria-current={item.current ? "page" : undefined}
												>
													<item.icon
														className={classNames(
															item.current
																? "text-white"
																: "text-indigo-300 group-hover:text-white",
															"mr-3 h-6 w-6"
														)}
														aria-hidden="true"
													/>
													<span>{item.name}</span>
												</Link>
											))}
										</div>
									</nav>
								</div>
							</Dialog.Panel>
						</Transition.Child>
						<div className="w-14 flex-shrink-0" aria-hidden="true">
							{/* Dummy element to force sidebar to shrink to fit close icon */}
						</div>
					</div>
				</Dialog>
			</Transition.Root>

			{/* Content area */}
			<div className="flex flex-1 flex-col overflow-hidden">
				<header className="w-full">
					<div className="relative z-10 flex h-16 flex-shrink-0 border-b border-gray-200 bg-white shadow-sm">
						<button
							type="button"
							className="border-r border-gray-200 px-4 text-gray-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500 md:hidden"
							onClick={() => setMobileMenuOpen(true)}
						>
							<span className="sr-only">Open sidebar</span>
							<Bars3BottomLeftIcon className="h-6 w-6" aria-hidden="true" />
						</button>
						<div className="flex flex-1 justify-between px-4 sm:px-6">
							<div className="flex flex-1">
								<form className="flex w-full md:ml-0" action="#" method="GET">
									<label htmlFor="desktop-search-field" className="sr-only">
										Search all files
									</label>
									<label htmlFor="mobile-search-field" className="sr-only">
										Search all files
									</label>
									<div className="relative w-full text-gray-400 focus-within:text-gray-600">
										<div className="pointer-events-none absolute inset-y-0 left-0 flex items-center">
											<MagnifyingGlassIcon
												className="h-5 w-5 flex-shrink-0"
												aria-hidden="true"
											/>
										</div>
										<input
											name="mobile-search-field"
											id="mobile-search-field"
											className="h-full w-full border-0 py-2 pl-8 pr-3 text-base text-gray-900 focus:outline-none focus:ring-0 focus:placeholder:text-gray-400 sm:hidden"
											placeholder="Search"
											type="search"
										/>
										<input
											name="desktop-search-field"
											id="desktop-search-field"
											className="hidden h-full w-full border-0 py-2 pl-8 pr-3 text-sm text-gray-900 focus:outline-none focus:ring-0 focus:placeholder:text-gray-400 sm:block"
											placeholder="Search all files"
											type="search"
										/>
									</div>
								</form>
							</div>
							<div className="ml-2 flex items-center space-x-4 sm:ml-6 sm:space-x-6">
								{/* Profile dropdown */}
								<Menu as="div" className="relative flex-shrink-0">
									<div>
										<Menu.Button className="flex rounded-full bg-white text-sm focus:outline-none focus:ring-2 focus:ring-indigo-600 focus:ring-offset-2">
											<span className="sr-only">Open user menu</span>
											<img
												className="h-8 w-8 rounded-full"
												src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=256&h=256&q=80"
												alt=""
											/>
										</Menu.Button>
									</div>
									<Transition
										as={Fragment}
										enter="transition ease-out duration-100"
										enterFrom="transform opacity-0 scale-95"
										enterTo="transform opacity-100 scale-100"
										leave="transition ease-in duration-75"
										leaveFrom="transform opacity-100 scale-100"
										leaveTo="transform opacity-0 scale-95"
									>
										<Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
											{userNavigation.map((item) => (
												<Menu.Item key={item.name}>
													{({ active }) => (
														<Link
															href={item.href}
															className={classNames(
																active ? "bg-gray-100" : "",
																"block px-4 py-2 text-sm text-gray-700"
															)}
														>
															{item.name}
														</Link>
													)}
												</Menu.Item>
											))}
										</Menu.Items>
									</Transition>
								</Menu>

								<button
									type="button"
									onClick={() => router.push("/upload")}
									className="rounded-full bg-indigo-600 p-1.5 text-white hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
								>
									<PlusIcon className="h-5 w-5" aria-hidden="true" />
									<span className="sr-only">Add file</span>
								</button>
							</div>
						</div>
					</div>
				</header>

				{children}
			</div>
		</div>
	);
}
