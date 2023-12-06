#!/usr/bin/env python

# with open("example", "r") as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.read().split("\n\n") if x.strip()]

# Define the almanac data
seeds = list(map(int, inp.pop(0).split(" ")[1:]))
seed_ranges = [seeds[i : i + 2] for i in range(0, len(seeds), 2)]

# Maps from the almanac
seed_to_soil_map = [[int(x) for x in y.split(" ")] for y in inp.pop(0).splitlines()[1:]]
soil_to_fertilizer_map = [
    [int(x) for x in y.split(" ")] for y in inp.pop(0).splitlines()[1:]
]
fertilizer_to_water_map = [
    [int(x) for x in y.split(" ")] for y in inp.pop(0).splitlines()[1:]
]
water_to_light_map = [
    [int(x) for x in y.split(" ")] for y in inp.pop(0).splitlines()[1:]
]
light_to_temperature_map = [
    [int(x) for x in y.split(" ")] for y in inp.pop(0).splitlines()[1:]
]
temperature_to_humidity_map = [
    [int(x) for x in y.split(" ")] for y in inp.pop(0).splitlines()[1:]
]
humidity_to_location_map = [
    [int(x) for x in y.split(" ")] for y in inp.pop(0).splitlines()[1:]
]

mappings = [
    seed_to_soil_map,
    soil_to_fertilizer_map,
    fertilizer_to_water_map,
    water_to_light_map,
    light_to_temperature_map,
    temperature_to_humidity_map,
    humidity_to_location_map,
]


def get_next_range(lookup, mapping):
    output = []
    # Since each seed is a start and a length
    for lookup_start, lookup_length in lookup:
        # get the mappings for the next part for the whole length
        while lookup_length > 0:
            for dest_range_start, src_range_start, length in mapping:
                delta = lookup_start - src_range_start
                # check if delta is what's contained in this part of the mapping
                if delta in range(length):
                    # only keep track of what this mapping maps to
                    length = min(length - delta, lookup_length)
                    output.append((dest_range_start + delta, length))
                    # truncate this part of the mapping
                    lookup_start += length
                    lookup_length -= length
                    break
            else:
                # if it's not in the range use the start and length provided
                output.append((lookup_start, lookup_length))
                break
    return output


for mapping in mappings:
    seed_ranges = get_next_range(seed_ranges, mapping)

min_locations = min(seed_ranges)[0]

print(min_locations)
