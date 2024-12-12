def is_monotone_within_range(levels):
    """
    Checks if the given list of levels is strictly monotone (either strictly increasing or strictly decreasing)
    with each adjacent difference between 1 and 3. Returns True if it's safe, False otherwise.
    """
    if len(levels) < 2:
        # A single level or empty list is trivially safe
        return True

    # Determine the trend (increasing or decreasing) based on the first non-equal pair
    trend = None  # 'inc' for increasing, 'dec' for decreasing
    for i in range(len(levels) - 1):
        diff = levels[i+1] - levels[i]
        if diff > 0:
            # Found a positive difference, so we must be strictly increasing
            if diff < 1 or diff > 3:
                return False
            trend = 'inc'
            break
        elif diff < 0:
            # Found a negative difference, so we must be strictly decreasing
            if (-diff) < 1 or (-diff) > 3:
                return False
            trend = 'dec'
            break
        # If diff == 0, not strictly monotone, fail immediately
        else:
            return False

    # If we never broke, it means all differences were zero (which isn't allowed)
    if trend is None:
        return False

    # Check the rest of the sequence for consistency
    for j in range(i+1, len(levels) - 1):
        diff = levels[j+1] - levels[j]
        if trend == 'inc':
            if diff <= 0 or diff > 3:
                return False
        else:  # trend == 'dec'
            if diff >= 0 or diff < -3:
                return False

    return True

def can_be_safe_with_one_removal(levels):
    # If already safe, no removal needed
    if is_monotone_within_range(levels):
        return True
    
    # Otherwise, try removing one level at a time
    for i in range(len(levels)):
        new_levels = levels[:i] + levels[i+1:]
        if is_monotone_within_range(new_levels):
            return True
    return False

def main():
    with open("input.txt", "r") as f:
        lines = f.read().strip().splitlines()

    count_safe = 0
    for line in lines:
        if not line.strip():
            continue
        levels = list(map(int, line.split()))
        
        if can_be_safe_with_one_removal(levels):
            count_safe += 1

    print(count_safe)

if __name__ == "__main__":
    main()
