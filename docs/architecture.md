# Sorting Engine Mechanics

The sorting logic uses a normalized chunk-sorting pipeline optimized for strict coordinate indexes.

### Operational Phases
1. **Coordinate Normalization**: A sorted proxy list assigns clean tracking ranks from 0 to N-1 to eliminate gaps between integers.
2. **Chunk Partition Pushing**: Elements matching narrow target index limits are sent to Stack B using a structured window size (15 for 100 entries, 35 for 500 entries).
3. **Optimized Restitution**: The highest tracking elements are targeted within Stack B and moved back to Stack A using the shortest rotation path available.