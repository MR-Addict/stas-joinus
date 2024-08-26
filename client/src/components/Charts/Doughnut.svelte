<script lang="ts">
	import * as d3 from 'd3';

	import colors from '$data/colors';
	import type { ChartDataType } from '$types/chart';

	export let title: string;
	export let data: ChartDataType[];

	let total = 0;
	let radius = 0;
	let innerSize = { width: 0, height: 0 };

	let pie: d3.Pie<any, ChartDataType>;
	let arc: d3.Arc<any, d3.PieArcDatum<ChartDataType>>;

	const config = {
		view: { width: 400, height: 400 },
		padding: { top: 10, right: 50, bottom: 0, left: 0 }
	};

	$: {
		total = data.reduce((acc, d) => acc + d.value, 0);

		innerSize.width = config.view.width - config.padding.left - config.padding.right;
		innerSize.height = config.view.height - config.padding.top - config.padding.bottom;
		radius = Math.min(innerSize.width, innerSize.height) / 2;

		arc = d3
			.arc<d3.PieArcDatum<ChartDataType>>()
			.innerRadius(radius * 0.5)
			.outerRadius(radius * 1);

		pie = d3
			.pie<ChartDataType>()
			.sort(null)
			.value((d) => d.value)
			.padAngle(0.01);
	}
</script>

<div class="w-full">
	<svg viewBox="0 0 {config.view.width} {config.view.height}" style="width: 100%; height:100%;">
		{#if total <= 0}
			<g transform="translate({config.view.width / 2},{radius + config.padding.top})">
				<circle r={radius} fill="lightgray" />

				<!-- animation circle -->
				{#key data}
					<circle r={Math.ceil(radius / 2 + 1)} fill="none" stroke="white" stroke-width={radius} class="animation" />
				{/key}
				<text text-anchor="middle" alignment-baseline="middle">暂无数据</text>
			</g>
		{:else}
			<g transform="translate({config.view.width / 2},{radius + config.padding.top})">
				{#each pie(data) as d, i (d.data.label)}
					<path d={arc(d)} fill={colors[i % colors.length]} />
					<text transform="translate({arc.centroid(d)})" text-anchor="middle" fill="white" alignment-baseline="middle">
						{d.data.value}
					</text>
				{/each}

				<!-- animation circle -->
				{#key data}
					<circle r={Math.ceil(radius / 2 + 1)} fill="none" stroke="white" stroke-width={radius} class="animation" />
				{/key}

				<!-- add total number in the center of doughnut -->
				{#key data}
					<text class="bar-value" text-anchor="middle" alignment-baseline="middle">
						{total}
					</text>
				{/key}
			</g>

			<!-- put lengends at the top right corner -->
			<g transform="translate({config.view.width / 2 + radius - config.padding.right}, {config.padding.top})">
				{#each data as d, i (d.label)}
					<g transform={`translate(0, ${i * 22})`}>
						<rect width="10" height="10" fill={colors[i % colors.length]} />
						<text x="15" dy=".4rem" font-size="0.85rem" alignment-baseline="middle">
							{`${d.label}(${((d.value / total) * 100).toFixed(0)}%)`}
						</text>
					</g>
				{/each}
			</g>
		{/if}

		<!-- add title under the doughnut -->
		<text
			x={config.view.width / 2}
			y={radius * 2 + config.padding.top + 20}
			text-anchor="middle"
			alignment-baseline="middle"
			font-weight="600"
		>
			{title}
		</text>
	</svg>
</div>

<style>
	.bar-value {
		opacity: 0;
		animation: scaelUp 0.5s forwards;
		animation-delay: 0.5s;
	}
	@keyframes scaelUp {
		to {
			opacity: 1;
		}
	}

	.animation {
		animation: draw 1s forwards;
	}

	@keyframes draw {
		from {
			stroke-dasharray: 1000, 0;
		}

		to {
			stroke-dasharray: 0, 1000;
		}
	}
</style>
