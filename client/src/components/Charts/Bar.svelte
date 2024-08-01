<script lang="ts">
	import * as d3 from 'd3';
	import colors from '$data/colors';
	import type { ChartDataType } from '$types/chart';

	export let title: string;
	export let data: ChartDataType[];
	export let size = { width: 800, height: 400 };

	let total = 0;
	let radius = 0;
	let innerSize = { width: 0, height: 0 };
	let x: d3.ScaleBand<string>;
	let y: d3.ScaleLinear<number, number>;

	const config = {
		minBarHeight: 5,
		padding: { top: 40, right: 0, bottom: 30, left: 0 }
	};

	$: {
		total = data.reduce((acc, d) => acc + d.value, 0);

		innerSize.width = size.width - config.padding.left - config.padding.right;
		innerSize.height = size.height - config.padding.top - config.padding.bottom;
		radius = Math.min(innerSize.width, innerSize.height) / 2;

		if (total > 0) {
			x = d3
				.scaleBand()
				.domain(data.map((d) => d.label))
				.range([0, innerSize.width])
				.padding(0.1);

			y = d3
				.scaleLinear()
				.domain([0, d3.max(data, (d) => d.value) || 0])
				.range([innerSize.height, 0]);
		}
	}
</script>

<svg width={size.width} height={size.height}>
	<text x={size.width / 2} text-anchor="middle" font-weight="600" alignment-baseline="before-edge">
		{title}
	</text>

	{#if total > 0}
		<g transform="translate({config.padding.left},{config.padding.top + 10})">
			{#each data as d, i (d.label)}
				<rect
					class="bar"
					x={x(d.label)}
					y={Math.min(y(d.value), innerSize.height - config.minBarHeight)}
					width={x.bandwidth()}
					height={Math.max(innerSize.height - y(d.value), config.minBarHeight)}
					fill={colors[i % colors.length]}
				/>
				<text
					x={(x(d.label) || 0) + x.bandwidth() / 2}
					y={Math.min(y(d.value), innerSize.height - config.minBarHeight) - 4}
					text-anchor="top"
					font-size="0.9rem"
				>
					{d.value}
				</text>
				<text
					x={(x(d.label) || 0) + x.bandwidth() / 2}
					y={innerSize.height}
					dy="0.2rem"
					text-anchor="middle"
					font-size="0.6rem"
					alignment-baseline="before-edge"
				>
					{d.label}
				</text>
			{/each}
		</g>
	{:else}
		{#key data}
			<rect
				rx={10}
				class="animate"
				fill="lightgray"
				width={innerSize.width}
				height={innerSize.height}
				transform="translate({config.padding.left},{config.padding.top})"
			/>
		{/key}

		<text
			text-anchor="middle"
			alignment-baseline="middle"
			transform="translate({innerSize.width / 2 + config.padding.left},{innerSize.height / 2 + config.padding.top})"
		>
			暂无数据
		</text>
	{/if}
</svg>

<style>
	.bar {
		opacity: 0.9;
		transition:
			height 0.5s ease,
			y 0.5s ease;
	}
	.animate {
		animation: fadeIn 2s;
	}
	@keyframes fadeIn {
		from {
			opacity: 0;
		}

		to {
			opacity: 1;
		}
	}
</style>
