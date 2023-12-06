#!/usr/bin/env python
import re

with open("input", "r") as f:
#with open("example", "r") as f:
    inp = [x.strip() for x in f.read().split('\n\n') if x.strip()]

seeds = inp

# Define the almanac data as provided in the example
seeds = list(map(int, inp.pop(0).split(' ')[1:]))

# Maps from the almanac
seed_to_soil_map = [[int(x) for x in y.split(' ')] for y in inp.pop(0).splitlines()[1:]]
soil_to_fertilizer_map = [[int(x) for x in y.split(' ')] for y in inp.pop(0).splitlines()[1:]]
fertilizer_to_water_map = [[int(x) for x in y.split(' ')] for y in inp.pop(0).splitlines()[1:]]
water_to_light_map = [[int(x) for x in y.split(' ')] for y in inp.pop(0).splitlines()[1:]]
light_to_temperature_map = [[int(x) for x in y.split(' ')] for y in inp.pop(0).splitlines()[1:]]
temperature_to_humidity_map = [[int(x) for x in y.split(' ')] for y in inp.pop(0).splitlines()[1:]]
humidity_to_location_map = [[int(x) for x in y.split(' ')] for y in inp.pop(0).splitlines()[1:]]

def find_mapped_value(source, mapping_data):
    for destination_start, source_start, length in mapping_data:
        if source_start <= source < source_start + length:
            return destination_start + (source - source_start)
    return source  # Return the same value if not mapped

# Function to process a seed through all the mappings to find its location
def process_seed_optimized(seed):
    soil = find_mapped_value(seed, seed_to_soil_map)
    fertilizer = find_mapped_value(soil, soil_to_fertilizer_map)
    water = find_mapped_value(fertilizer, fertilizer_to_water_map)
    light = find_mapped_value(water, water_to_light_map)
    temperature = find_mapped_value(light, light_to_temperature_map)
    humidity = find_mapped_value(temperature, temperature_to_humidity_map)
    location = find_mapped_value(humidity, humidity_to_location_map)
    return location

# Process each seed and find the lowest location number using the optimized approach
locations_optimized = [process_seed_optimized(seed) for seed in seeds]
lowest_location_optimized = min(locations_optimized)
print(lowest_location_optimized)
