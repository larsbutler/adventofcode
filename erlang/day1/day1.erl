-module(day1).

-export([main/1]).


solve(Input) ->
    % for each input
    % [First, Rest] = Input,
    StartAngle = half_pi(),
    StartXLoc = 0,
    StartYLoc = 0,
    {XLoc, YLoc, _Angle} = solve(Input, StartXLoc, StartYLoc, StartAngle),
    Solution = round(abs(XLoc) + abs(YLoc)),
    Solution.

% base case for solve/4
solve([], XLoc, YLoc, Angle) ->
    {XLoc, YLoc, Angle};
solve([Head | Tail], StartXLoc, StartYLoc, StartAngle) ->
    {XLoc, YLoc, Angle} = solve_one(Head, StartXLoc, StartYLoc, StartAngle),
    solve(Tail, XLoc, YLoc, Angle).

solve_one([Direction | Distance], XLoc, YLoc, Angle) ->
    % Convert the distance string to an int
    IntDistance = list_to_integer(Distance),

    Sign = math:pow(-1, 1 + ((Direction rem 10) + (Direction div 10) rem 2)),
    NewAngle = Angle + Sign * half_pi(),

    X = math:cos(NewAngle),
    Y = math:sin(NewAngle),
    VX = X * IntDistance,
    VY = Y * IntDistance,

    NewXLoc = XLoc + VX,
    NewYLoc = YLoc + VY,
    {NewXLoc, NewYLoc, NewAngle}.


process_input(InputFile) ->
    {ok, File} = file:read_file(InputFile),
    Content = unicode:characters_to_list(File),

    [Puzzle | Expected] = string:tokens(Content, "\n"),

    Input = string:tokens(Puzzle, ", "),
    Solution = solve(Input),

    case Expected of
        "" ->
            io:format("Solution: ~w.~n", [Solution]);
        _Else ->
            io:format("Solution: ~w. Expected: ~s.~n", [Solution, Expected])
    end.

half_pi() ->
    math:pi() / 2.


main(Args) ->
    case Args of
        [InputFile | _ ] ->
            process_input(InputFile);
        [] ->
            ScriptName = escript:script_name(),
            io:format("Usage: $ escript ~s INPUT_FILE~n", [ScriptName]),
            halt(1)
    end.
