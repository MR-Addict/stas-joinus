<script lang="ts">
	import * as d3 from 'd3';
	import { draw } from 'svelte/transition';

	import colors from '$data/colors';

	type DoughnutDataType = { label: string; value: number };

	export let title: string;
	export let data: DoughnutDataType[] = [];
	export let size = { width: 380, height: 370 };
	export let margin = { top: 0, right: 0, bottom: 0, left: 0 };

	let total = 0;
	let radius = 0;
	let innerSize = { width: 0, height: 0 };

	let pie: d3.Pie<any, DoughnutDataType>;
	let arc: d3.Arc<any, d3.PieArcDatum<DoughnutDataType>>;

	const config = {
		textAnchor: 'middle',
		padding: { top: 5, right: 50, bottom: 30, left: 0 }
	};

	$: {
		total = data.reduce((acc, d) => acc + d.value, 0);

		innerSize.width = size.width - margin.left - margin.right - config.padding.left - config.padding.right;
		innerSize.height = size.height - margin.top - margin.bottom - config.padding.top - config.padding.bottom;
		radius = Math.min(innerSize.width, innerSize.height) / 2;

		arc = d3
			.arc<d3.PieArcDatum<DoughnutDataType>>()
			.innerRadius(radius * 0.5)
			.outerRadius(radius * 1);

		pie = d3
			.pie<DoughnutDataType>()
			.sort((a, b) => b.label.localeCompare(a.label))
			.value((d) => d.value);
	}
</script>

<svg width={size.width} height={size.height}>
	{#if total <= 0}
		<circle cx={size.width / 2} cy={radius + config.padding.top} r={radius} fill="lightgray" />
		<text x={size.width / 2} y={size.height / 2} text-anchor={config.textAnchor}>暂无数据</text>
		<text x={size.width / 2} y={radius * 2 + config.padding.top + 20} text-anchor={config.textAnchor} font-weight="600">
			{title}
		</text>
	{:else}
		<g transform="translate({size.width / 2},{radius + config.padding.top})">
			{#each pie(data.filter((d) => d.value > 0)) as d, i (d.data.label)}
				<path d={arc(d)} fill={colors[i % colors.length]} out:draw={{ duration: 300 }} />
				<g transform="translate({arc.centroid(d)})">
					<text text-anchor={config.textAnchor} dy="0" fill="white">
						{d.data.value}
					</text>
				</g>
			{/each}

			<!-- add total number in the center of doughnut -->
			<text x="0" y="0" text-anchor={config.textAnchor} dy=".45rem">
				{total}
			</text>
		</g>

		<!-- add title under the doughnut -->
		<text x={size.width / 2} y={radius * 2 + config.padding.top + 20} text-anchor={config.textAnchor} font-weight="600">
			{title}
		</text>

		<!-- put lengends at the top right corner -->
		<g transform="translate({size.width / 2 + radius - config.padding.right}, {config.padding.top})">
			{#each data as d, i (d.label)}
				<g transform={`translate(0, ${i * 22})`}>
					<rect width="12" height="12" fill={colors[i % colors.length]} />
					<text x="17" y="5" dy=".35em" font-size="0.85rem">
						{`${d.label}(${((d.value / total) * 100).toFixed(0)}%)`}
					</text>
				</g>
			{/each}
		</g>
	{/if}
</svg>
