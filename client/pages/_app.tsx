import "@/styles/globals.css";
import type { AppProps } from "next/app";
import { Toaster } from "react-hot-toast";

import { SettingsProvider } from "@/context/SettingsProvider";

export default function App({ Component, pageProps }: AppProps) {
	return (
		<>
			<SettingsProvider>
				<Toaster />
				<Component {...pageProps} />
			</SettingsProvider>
		</>
	);
}
