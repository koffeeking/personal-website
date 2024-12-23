package templa

import (
	"github.com/go-fuego/fuego/examples/full-app-gourmet/store"
	"github.com/go-fuego/fuego/examples/full-app-gourmet/templa/components"
)

type HomeProps struct {
	Recipes        []store.Recipe
	PopularRecipes []store.Recipe
	FastRecipes    []store.Recipe
	HealthyRecipes []store.Recipe
}

 Home(props HomeProps) {
	@page("Gourmet") {
		<section
			id="hero"
      class="flex w-full flex-col items-center justify-center gap-4 bg-gray-800 bg-opacity-50 bg-cover bg-top bg-no-repeat px-4 py-8 text-center  text-white bg-blend-color-burn sm:px-6 md:gap-8 md:px-8 md:py-16 lg:px-10 md:my-8 md:rounded-xl"
			style="background-image: url('/static/hero.webp')"
		>
			<div class="backdrop-blur bg-white/10 p-4 px-8 rounded-lg">
				<h1 class="font-serif text-4xl font-bold italic text-white">Gourmet</h1>
				<p class="text-lg italic">Healthy recipes</p>
			</div>
		</section>
		<section
			id="recipes-section"
			class="flex w-full flex-col gap-4 p-2 md:gap-8 md:p-4"
		>
			<h2>Today's recipes 📅</h2>
			@components.SliderRecipes(props.Recipes)
		</section>
		<section
			id="popular-recipes-section"
			class="flex w-full flex-col gap-4 p-2 md:gap-8 md:p-4"
		>
			<h2>Popular recipes this week ⭐️</h2>
			@components.SliderRecipes(props.PopularRecipes)
		</section>
		<section
			id="fast-recipes-section"
			class="flex w-full flex-col gap-4 p-2 md:gap-8 md:p-4"
		>
			<h2>Fast recipes ⚡️</h2>
			@components.SliderRecipes(props.FastRecipes)
		</section>
		<section
			id="healthy-recipes-section"
			class="flex w-full flex-col gap-4 p-2 md:gap-8 md:p-4"
		>
			<h2>Healthy recipes 🥗</h2>
			@components.SliderRecipes(props.HealthyRecipes)
		</section>
	}
}
