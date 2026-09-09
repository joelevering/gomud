#!/usr/bin/env python3
"""Consistency check for a rooms.json-shaped file. Exit 0 if clean, 1 if not."""
import json
import sys

if len(sys.argv) != 2:
    print("usage: validate_rooms.py <path-to-rooms-json>")
    sys.exit(2)

rooms = json.load(open(sys.argv[1]))
ids = [r["id"] for r in rooms]
id_set = set(ids)
problems = []

if len(id_set) != len(ids):
    seen = set()
    dupes = sorted({i for i in ids if i in seen or seen.add(i)})
    problems.append(f"duplicate room ids: {dupes}")

dangling = []
for r in rooms:
    for e in r.get("exits", []):
        if e["room_id"] not in id_set:
            dangling.append((r["id"], e["room_id"]))
if dangling:
    problems.append(f"dangling exit targets (from_room -> missing_target): {dangling}")

npc_ids = []
for r in rooms:
    for n in r.get("npcs", []):
        npc_ids.append(n["id"])
if len(set(npc_ids)) != len(npc_ids):
    seen = set()
    dupes = sorted({i for i in npc_ids if i in seen or seen.add(i)})
    problems.append(f"duplicate npc ids: {dupes}")

areas = {}
for r in rooms:
    area = r["name"].split(" - ")[0].strip()
    areas[area] = areas.get(area, 0) + 1

print(f"rooms: {len(rooms)}  npcs: {len(npc_ids)}  areas: {len(areas)}")
for area, count in sorted(areas.items()):
    print(f"  {area}: {count} rooms")

if problems:
    print("\nPROBLEMS FOUND:")
    for p in problems:
        print(f"  - {p}")
    sys.exit(1)

print("\nno problems found")
sys.exit(0)
