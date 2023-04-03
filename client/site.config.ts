import { siteConfig } from "@/types/site-config";
import {
	BookOpenIcon,
	DocumentDuplicateIcon,
	HomeIcon,
	UserCircleIcon,
} from "@heroicons/react/24/outline";

import auvidLogo from "./public/auvid.svg";

const domainName = process.env.DOMAIN_NAME || "app.auvid.io";

export default siteConfig({
	siteName: "HEIC Image Converter",
	productBrand: auvidLogo,
	domain: domainName,
	developer: "Alchemized Software Ltd.",
	contactEmail: "hello@alchemizedsoftware.com",

	description: "A really good HEIC image converter.",

	// main navigation tabs
	mainNavTabs: [
		{ name: "Dashboard", icon: HomeIcon, current: true, href: "/" },
		{
			name: "Notes",
			icon: DocumentDuplicateIcon,
			current: false,
			children: [
				{ name: "Overview", href: "/notes" },
				{ name: "Add Note", href: "/upload" },
			],
		},
		{
			name: "Account",
			icon: UserCircleIcon,
			current: false,
			children: [
				{ name: "Your Profile", href: "/account" },
				// TODO: hide until we have membership support
				// { name: 'Subscription', href: '/account/subscription' },
				{
					name: "About Avuid",
					icon: BookOpenIcon,
					current: false,
					href: "/about",
				},
			],
		},
	],
});
