<template>
	<div class="flex flex-col items-center">
		<h1 class="text-gray-800 text-5xl font-bold">YOUTUBE-MP3</h1>
		<transition
			enter-active-class="animate__animated animate__fadeIn"
			leave-active-class="animate__animated animate__fadeOut"
			v-on:after-leave="done = true"
			appear
		>
			<div
				class="container mx-auto flex flex-col items-center max-w-4xl"
			>
				<form
					class="shadow-xl mt-5 w-80 md:w-8/12 p-4 flex flex-col bg-white rounded-lg space-y-5 md:space-y-8"
					@submit.prevent
				>
					<label
						for="url"
						class="font-semibold text-xl"
						>Enter a YouTube URL to
						convert</label
					>
					<input
						ref="firstinput"
						v-model="url"
						name="url"
						type="url"
						autocomplete="off"
						placeholder="https://youtu.be/vhCEaGWTu10"
						class="mb-3 py-3 px-4 border border-gray-400 focus:outline-none rounded-md focus:ring-1 ring-cyan-500"
					/>
					{{ this.tempname }}
					<div class="">
						<input
							autocomplete="off"
							name="songname"
							type="text"
							v-model="songname"
							placeholder="Enter name of the song"
							class="mb-3 py-3 px-4 border border-gray-400 focus:outline-none rounded-md focus:ring-1 ring-cyan-500"
						/>
					</div>
					<button
						@click="convert()"
						class="w-full bg-blue-500 text-white p-3 rounded-lg font-semibold text-lg"
						:class="{'animate-pulse': !done}"
						:disabled="!done"
					>
						Convert
					</button>
				</form>
			</div>
		</transition>
	</div>
</template>

<script>
import { download } from "./download";
/* let download = require('./download') */
export default {
	name: "UrlForm",
	data() {
		return {
			url: "",
			done: true,
			received: false,
			songname: "",
			tempname: "",
			checkurl: "http://localhost:3001/song/",
			inter: null,
		};
	},
	methods: {
		convert: function () {
			let re = /((([A-Za-z]{3,9}:(?:\/\/)?)(?:[\-;:&=\+\$,\w]+@)?[A-Za-z0-9\.\-]+|(?:www\.|[\-;:&=\+\$,\w]+@)[A-Za-z0-9\.\-]+)((?:\/[\+~%\/\.\w\-_]*)?\??(?:[\-\+=&;%@\.\w_]*)#?(?:[\.\!\/\\\w]*))?)/;

			if (!re.test(this.url)) {
				this.$toast.error(`Enter a valid URL`, {
					position: "top",
				});
				return;
			}
			this.done = false;
			this.received = false;
			fetch("http://localhost:3001", {
				method: "POST",
				headers: {
					Accept:
						"application/json, text,plain, */*",
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					url: this.url,
				}),
			})
				.then((res) => res.text())
				.then((text) => {
					/* this.done = true; */
					this.received = false;
					/* console.log(text) */
					this.tempname = text;
					this.check();
				});
		},
		check: function () {
			var ref = this;
			this.inter = setInterval(function () {
				fetch(
					`http://localhost:3001/song/${ref.tempname}`
				)
					.then((res) => {
						/* console.log(ref.tempname); */
						if (res.status == 200) {
							ref.received = true;
							ref.tempname = "";
							ref.clearinter();
							return res.blob();
						}
					})
					.then((blob) => {
						if (ref.received == true) {
							ref.done = true;
							download(
								blob,
								ref.songname
							);
						}
					});
			}, 500);
		},
		clearinter: function () {
			clearInterval(this.inter);
		},
	},
};
</script>
