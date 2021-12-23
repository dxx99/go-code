package main


var exampleData = []byte(`[{
    "location": {
        "x": 407838351,
        "y": -746143763
    },
    "name": "Patriots Path, Mendham, NJ 07945, USA"
}, {
    "location": {
        "x": 408122808,
        "y": -743999179
    },
    "name": "101 New Jersey 10, Whippany, NJ 07981, USA"
}, {
    "location": {
        "x": 413628156,
        "y": -749015468
    },
    "name": "U.S. 6, Shohola, PA 18458, USA"
}, {
    "location": {
        "x": 419999544,
        "y": -740371136
    },
    "name": "5 Conners Road, Kingston, NY 12401, USA"
}, {
    "location": {
        "x": 414008389,
        "y": -743951297
    },
    "name": "Mid Hudson Psychiatric Center, New Hampton, NY 10958, USA"
}, {
    "location": {
        "x": 419611318,
        "y": -746524769
    },
    "name": "287 Flugertown Road, Livingston Manor, NY 12758, USA"
}, {
    "location": {
        "x": 406109563,
        "y": -742186778
    },
    "name": "4001 Tremley Point Road, Linden, NJ 07036, USA"
}, {
    "location": {
        "x": 416802456,
        "y": -742370183
    },
    "name": "352 South Mountain Road, Wallkill, NY 12589, USA"
}, {
    "location": {
        "x": 412950425,
        "y": -741077389
    },
    "name": "Bailey Turn Road, Harriman, NY 10926, USA"
}, {
    "location": {
        "x": 412144655,
        "y": -743949739
    },
    "name": "193-199 Wawayanda Road, Hewitt, NJ 07421, USA"
}, {
    "location": {
        "x": 415736605,
        "y": -742847522
    },
    "name": "406-496 Ward Avenue, Pine Bush, NY 12566, USA"
}, {
    "location": {
        "x": 413843930,
        "y": -740501726
    },
    "name": "162 Merrill Road, Highland Mills, NY 10930, USA"
}, {
    "location": {
        "x": 410873075,
        "y": -744459023
    },
    "name": "Clinton Road, West Milford, NJ 07480, USA"
}, {
    "location": {
        "x": 412346009,
        "y": -744026814
    },
    "name": "16 Old Brook Lane, Warwick, NY 10990, USA"
}, {
    "location": {
        "x": 402948455,
        "y": -747903913
    },
    "name": "3 Drake Lane, Pennington, NJ 08534, USA"
}, {
    "location": {
        "x": 406337092,
        "y": -740122226
    },
    "name": "6324 8th Avenue, Brooklyn, NY 11220, USA"
}, {
    "location": {
        "x": 406421967,
        "y": -747727624
    },
    "name": "1 Merck Access Road, Whitehouse Station, NJ 08889, USA"
}, {
    "location": {
        "x": 416318082,
        "y": -749677716
    },
    "name": "78-98 Schalck Road, Narrowsburg, NY 12764, USA"
}, {
    "location": {
        "x": 415301720,
        "y": -748416257
    },
    "name": "282 Lakeview Drive Road, Highland Lake, NY 12743, USA"
}, {
    "location": {
        "x": 402647019,
        "y": -747071791
    },
    "name": "330 Evelyn Avenue, Hamilton Township, NJ 08619, USA"
}, {
    "location": {
        "x": 412567807,
        "y": -741058078
    },
    "name": "New York State Reference Route 987E, Southfields, NY 10975, USA"
}, {
    "location": {
        "x": 416855156,
        "y": -744420597
    },
    "name": "103-271 Tempaloni Road, Ellenville, NY 12428, USA"
}, {
    "location": {
        "x": 404663628,
        "y": -744820157
    },
    "name": "1300 Airport Road, North Brunswick Township, NJ 08902, USA"
}, {
    "location": {
        "x": 407113723,
        "y": -749746483
    },
    "name": ""
}, {
    "location": {
        "x": 402133926,
        "y": -743613249
    },
    "name": ""
}, {
    "location": {
        "x": 400273442,
        "y": -741220915
    },
    "name": ""
}, {
    "location": {
        "x": 411236786,
        "y": -744070769
    },
    "name": ""
}, {
    "location": {
        "x": 411633782,
        "y": -746784970
    },
    "name": "211-225 Plains Road, Augusta, NJ 07822, USA"
}, {
    "location": {
        "x": 415830701,
        "y": -742952812
    },
    "name": ""
}, {
    "location": {
        "x": 413447164,
        "y": -748712898
    },
    "name": "165 Pedersen Ridge Road, Milford, PA 18337, USA"
}, {
    "location": {
        "x": 405047245,
        "y": -749800722
    },
    "name": "100-122 Locktown Road, Frenchtown, NJ 08825, USA"
}, {
    "location": {
        "x": 418858923,
        "y": -746156790
    },
    "name": ""
}, {
    "location": {
        "x": 417951888,
        "y": -748484944
    },
    "name": "650-652 Willi Hill Road, Swan Lake, NY 12783, USA"
}, {
    "location": {
        "x": 407033786,
        "y": -743977337
    },
    "name": "26 East 3rd Street, New Providence, NJ 07974, USA"
}, {
    "location": {
        "x": 417548014,
        "y": -740075041
    },
    "name": ""
}, {
    "location": {
        "x": 410395868,
        "y": -744972325
    },
    "name": ""
}, {
    "location": {
        "x": 404615353,
        "y": -745129803
    },
    "name": ""
}, {
    "location": {
        "x": 406589790,
        "y": -743560121
    },
    "name": "611 Lawrence Avenue, Westfield, NJ 07090, USA"
}, {
    "location": {
        "x": 414653148,
        "y": -740477477
    },
    "name": "18 Lannis Avenue, New Windsor, NY 12553, USA"
}, {
    "location": {
        "x": 405957808,
        "y": -743255336
    },
    "name": "82-104 Amherst Avenue, Colonia, NJ 07067, USA"
}, {
    "location": {
        "x": 411733589,
        "y": -741648093
    },
    "name": "170 Seven Lakes Drive, Sloatsburg, NY 10974, USA"
}, {
    "location": {
        "x": 412676291,
        "y": -742606606
    },
    "name": "1270 Lakes Road, Monroe, NY 10950, USA"
}, {
    "location": {
        "x": 409224445,
        "y": -748286738
    },
    "name": "509-535 Alphano Road, Great Meadows, NJ 07838, USA"
}, {
    "location": {
        "x": 406523420,
        "y": -742135517
    },
    "name": "652 Garden Street, Elizabeth, NJ 07202, USA"
}, {
    "location": {
        "x": 401827388,
        "y": -740294537
    },
    "name": "349 Sea Spray Court, Neptune City, NJ 07753, USA"
}, {
    "location": {
        "x": 410564152,
        "y": -743685054
    },
    "name": "13-17 Stanley Street, West Milford, NJ 07480, USA"
}, {
    "location": {
        "x": 408472324,
        "y": -740726046
    },
    "name": "47 Industrial Avenue, Teterboro, NJ 07608, USA"
}, {
    "location": {
        "x": 412452168,
        "y": -740214052
    },
    "name": "5 White Oak Lane, Stony Point, NY 10980, USA"
}, {
    "location": {
        "x": 409146138,
        "y": -746188906
    },
    "name": "Berkshire Valley Management Area Trail, Jefferson, NJ, USA"
}, {
    "location": {
        "x": 404701380,
        "y": -744781745
    },
    "name": "1007 Jersey Avenue, New Brunswick, NJ 08901, USA"
}, {
    "location": {
        "x": 409642566,
        "y": -746017679
    },
    "name": "6 East Emerald Isle Drive, Lake Hopatcong, NJ 07849, USA"
}, {
    "location": {
        "x": 408031728,
        "y": -748645385
    },
    "name": "1358-1474 New Jersey 57, Port Murray, NJ 07865, USA"
}, {
    "location": {
        "x": 413700272,
        "y": -742135189
    },
    "name": "367 Prospect Road, Chester, NY 10918, USA"
}, {
    "location": {
        "x": 404310607,
        "y": -740282632
    },
    "name": "10 Simon Lake Drive, Atlantic Highlands, NJ 07716, USA"
}, {
    "location": {
        "x": 409319800,
        "y": -746201391
    },
    "name": "11 Ward Street, Mount Arlington, NJ 07856, USA"
}, {
    "location": {
        "x": 406685311,
        "y": -742108603
    },
    "name": "300-398 Jefferson Avenue, Elizabeth, NJ 07201, USA"
}, {
    "location": {
        "x": 419018117,
        "y": -749142781
    },
    "name": "43 Dreher Road, Roscoe, NY 12776, USA"
}, {
    "location": {
        "x": 412856162,
        "y": -745148837
    },
    "name": "Swan Street, Pine Island, NY 10969, USA"
}, {
    "location": {
        "x": 416560744,
        "y": -746721964
    },
    "name": "66 Pleasantview Avenue, Monticello, NY 12701, USA"
}, {
    "location": {
        "x": 405314270,
        "y": -749836354
    },
    "name": ""
}, {
    "location": {
        "x": 414219548,
        "y": -743327440
    },
    "name": ""
}, {
    "location": {
        "x": 415534177,
        "y": -742900616
    },
    "name": "565 Winding Hills Road, Montgomery, NY 12549, USA"
}, {
    "location": {
        "x": 406898530,
        "y": -749127080
    },
    "name": "231 Rocky Run Road, Glen Gardner, NJ 08826, USA"
}, {
    "location": {
        "x": 407586880,
        "y": -741670168
    },
    "name": "100 Mount Pleasant Avenue, Newark, NJ 07104, USA"
}, {
    "location": {
        "x": 400106455,
        "y": -742870190
    },
    "name": "517-521 Huntington Drive, Manchester Township, NJ 08759, USA"
}, {
    "location": {
        "x": 400066188,
        "y": -746793294
    },
    "name": ""
}, {
    "location": {
        "x": 418803880,
        "y": -744102673
    },
    "name": "40 Mountain Road, Napanoch, NY 12458, USA"
}, {
    "location": {
        "x": 414204288,
        "y": -747895140
    },
    "name": ""
}, {
    "location": {
        "x": 414777405,
        "y": -740615601
    },
    "name": ""
}, {
    "location": {
        "x": 415464475,
        "y": -747175374
    },
    "name": "48 North Road, Forestburgh, NY 12777, USA"
}, {
    "location": {
        "x": 404062378,
        "y": -746376177
    },
    "name": ""
}, {
    "location": {
        "x": 405688272,
        "y": -749285130
    },
    "name": ""
}, {
    "location": {
        "x": 400342070,
        "y": -748788996
    },
    "name": ""
}, {
    "location": {
        "x": 401809022,
        "y": -744157964
    },
    "name": ""
}, {
    "location": {
        "x": 404226644,
        "y": -740517141
    },
    "name": "9 Thompson Avenue, Leonardo, NJ 07737, USA"
}, {
    "location": {
        "x": 410322033,
        "y": -747871659
    },
    "name": ""
}, {
    "location": {
        "x": 407100674,
        "y": -747742727
    },
    "name": ""
}, {
    "location": {
        "x": 418811433,
        "y": -741718005
    },
    "name": "213 Bush Road, Stone Ridge, NY 12484, USA"
}, {
    "location": {
        "x": 415034302,
        "y": -743850945
    },
    "name": ""
}, {
    "location": {
        "x": 411349992,
        "y": -743694161
    },
    "name": ""
}, {
    "location": {
        "x": 404839914,
        "y": -744759616
    },
    "name": "1-17 Bergen Court, New Brunswick, NJ 08901, USA"
}, {
    "location": {
        "x": 414638017,
        "y": -745957854
    },
    "name": "35 Oakland Valley Road, Cuddebackville, NY 12729, USA"
}, {
    "location": {
        "x": 412127800,
        "y": -740173578
    },
    "name": ""
}, {
    "location": {
        "x": 401263460,
        "y": -747964303
    },
    "name": ""
}, {
    "location": {
        "x": 412843391,
        "y": -749086026
    },
    "name": ""
}, {
    "location": {
        "x": 418512773,
        "y": -743067823
    },
    "name": ""
}, {
    "location": {
        "x": 404318328,
        "y": -740835638
    },
    "name": "42-102 Main Street, Belford, NJ 07718, USA"
}, {
    "location": {
        "x": 419020746,
        "y": -741172328
    },
    "name": ""
}, {
    "location": {
        "x": 404080723,
        "y": -746119569
    },
    "name": ""
}, {
    "location": {
        "x": 401012643,
        "y": -744035134
    },
    "name": ""
}, {
    "location": {
        "x": 404306372,
        "y": -741079661
    },
    "name": ""
}, {
    "location": {
        "x": 403966326,
        "y": -748519297
    },
    "name": ""
}, {
    "location": {
        "x": 405002031,
        "y": -748407866
    },
    "name": ""
}, {
    "location": {
        "x": 409532885,
        "y": -742200683
    },
    "name": ""
}, {
    "location": {
        "x": 416851321,
        "y": -742674555
    },
    "name": ""
}, {
    "location": {
        "x": 406411633,
        "y": -741722051
    },
    "name": "3387 Richmond Terrace, Staten Island, NY 10303, USA"
}, {
    "location": {
        "x": 413069058,
        "y": -744597778
    },
    "name": "261 Van Sickle Road, Goshen, NY 10924, USA"
}, {
    "location": {
        "x": 418465462,
        "y": -746859398
    },
    "name": ""
}, {
    "location": {
        "x": 411733222,
        "y": -744228360
    },
    "name": ""
}, {
    "location": {
        "x": 410248224,
        "y": -747127767
    },
    "name": "3 Hasta Way, Newton, NJ 07860, USA"
}]`)

