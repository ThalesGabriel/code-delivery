import { Box, Button, FormControl, Grid, InputLabel, makeStyles, MenuItem, Select, Theme } from '@material-ui/core';
import { Loader } from 'google-maps';
import { sample, shuffle } from 'lodash';
import { useSnackbar } from 'notistack';
import { FormEvent, MouseEvent, useCallback, useEffect, useRef, useState } from 'react';
import { getCurrentPosition } from '../util/geolocation';
import { makeCarIcon, makeMarkerIcon, Map } from '../util/map';
import { Route } from '../util/models';
import { Navbar } from './Navbar';
import io from "socket.io-client";
import { RouteExistsError } from '../errors/route-exists.error';

const colors = [
  "#b71c1c",
  "#4a148c",
  "#2e7d32",
  "#e65100",
  "#2962ff",
  "#c2185b",
  "#FFCD00",
  "#3e2723",
  "#03a9f4",
  "#827717",
];

const useStyles = makeStyles((theme: Theme) => ({
	formControl: {
		margin: theme.spacing(1),
		minWidth: 120,
		width: '100%'
	},
	selectEmpty: {
		marginTop: theme.spacing(2),
	},
	root: {
    width: "100%",
    height: "100%",
  },
  form: {
    margin: "16px",
  },
  btnSubmitWrapper: {
    textAlign: "center",
    marginTop: "8px",
  },
  map: {
    width: "100%",
    height: "100%",
  },
}));

const googleMapsLoader = new Loader(process.env.REACT_APP_GOOGLE_API_KEY)
const API_URL = process.env.REACT_APP_API_URL as string;

export const Mapping = () => {
	const classes = useStyles();
	const [routes, setRoutes] = useState<Route[]>([]);
	const [routeIdSelected, setRouteIdSelected] = useState<string>("");
	const mapRef = useRef<Map>();
	const socketIORef = useRef<SocketIOClient.Socket>();
	const { enqueueSnackbar } = useSnackbar();

	const finishRoute = useCallback(
    (route: Route) => {
      enqueueSnackbar(`${route.title} finalizou!`, {
        variant: "success",
      });
      mapRef.current?.removeRoute(route._id);
    },
    [enqueueSnackbar]
  );

	useEffect(() => {
    if (!socketIORef.current?.connected) {
      socketIORef.current = io.connect(API_URL);
      socketIORef.current.on("connect", () => console.log("conectou"));
    }

    const handler = (data: {
      routeId: string;
      positions: [number, number];
      finished: boolean;
    }) => {
      console.log(data);
      mapRef.current?.moveCurrentMarker(data.routeId, {
        lat: data.positions[0],
        lng: data.positions[1],
      });
      const route = routes.find((route) => route._id === data.routeId) as Route;
      if (data.finished) {
        finishRoute(route);
      }
    };
    socketIORef.current?.on("new-position", handler);
    return () => {
      socketIORef.current?.off("new-position", handler);
    };
  }, [finishRoute, routes, routeIdSelected]);

  const startRoute = useCallback(
    (event: FormEvent) => {
      event.preventDefault();
      const route = routes.find((route) => route._id === routeIdSelected);
      const color = sample(shuffle(colors)) as string;
      try {
        mapRef.current?.addRoute(routeIdSelected, {
          currentMarkerOptions: {
            position: route?.startPosition,
            icon: makeCarIcon(color),
          },
          endMarkerOptions: {
            position: route?.endPosition,
            icon: makeMarkerIcon(color),
          },
        });
        socketIORef.current?.emit("new-direction", {
          routeId: routeIdSelected,
        });
      } catch (error) {
        if (error instanceof RouteExistsError) {
          enqueueSnackbar(`${route?.title} já adicionado, espere finalizar.`, {
            variant: "error",
          });
          return;
        }
        throw error;
      }
    },
    [routeIdSelected, routes, enqueueSnackbar]
  );

	useEffect(() => {
		fetch(`${API_URL}/routes`)
			.then(data => data.json())
			.then((data) => setRoutes(data))
	}, [ ])

	useEffect(() => {
		(async () => {
			const [, position] = await Promise.all([
				googleMapsLoader.load(),
				getCurrentPosition({enableHighAccuracy: true})
			]);
			const divMap = document.getElementById("map") as HTMLElement;
			mapRef.current = new Map(divMap, {
        zoom: 15,
        center: position,
      });
		})()
	}, [])

	return (
		<Grid container style={{ height: '100%', width: '100%' }}>
			<Grid item container xs={12} sm={3} style={{ backgroundColor: 'black', height: 'auto' }}>
				<Box style={{ width: '100%'}}>
					<Box style={{ textAlign: 'center'}}>
						<Navbar />
					</Box>
					<Box style={{ width: '100%', padding: 10 }}>
						<Grid item xs={11} style={{marginTop: 20}}>
							<form onSubmit={startRoute}>
							<FormControl variant="filled" className={classes.formControl}>
									<InputLabel htmlFor="filled-age-native-simple">Selecione uma rota</InputLabel>
									<Select
										native
										value={routeIdSelected}
										onChange={(event) => setRouteIdSelected(event.target.value + "")}
										inputProps={{
											name: 'route',
											id: 'filled-age-native-simple',
										}}
									>
										<option aria-label="None" value="" />
										{routes.map((route, index) => (
											<>
												<option key={index} value={route._id}>{route.title}</option>
											</>
										))}
									</Select>
								</FormControl>
								<Box style={{textAlign: 'center', marginTop: 20}}>
									<Button color="primary" type="submit" variant="contained">Iniciar</Button>
								</Box>
							</form>
						</Grid>
					</Box>
				</Box>
			</Grid>
			<Grid item container xs={12} sm={9} >
				<div id="map" style={{height: '100%', width: '100%'}}></div>
			</Grid>
		</Grid>
	);
}
