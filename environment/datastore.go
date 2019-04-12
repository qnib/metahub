package environment

/*
	manifestClient "github.com/docker/distribution/registry/client"

	datastoreClient, err := datastore.NewClient(ctx, "")
	if err != nil {
		log.Printf("failed to create client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	manifestEntityKey := datastore.NameKey(entityKindManifest, string(newDigest), nil)
	me := manifestEntity{
		Content: []byte{},
	}
	if _, err := datastoreClient.Put(ctx, manifestEntityKey, &me); err != nil {
		log.Printf("error saving new manifest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

			manifestEntityKey := datastore.NameKey(entityKindManifest, string(d), nil)
		me := manifestEntity{}
		if err := datastoreClient.Get(ctx, manifestEntityKey, &me); err != nil && err != datastore.ErrNoSuchEntity {
			log.Printf("error looking existing manifest: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if err == datastore.ErrNoSuchEntity {
			// TODO think ... required??
			log.Printf("THIS SHOULD NOT HAPPEN")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			newNanifest := &manifestSchema2.DeserializedManifest{}
			if err := newNanifest.UnmarshalJSON(me.Content); err != nil {
				log.Printf("error parsing manifest (%d bytes, key=%v) backing in database: %v", len(me.Content), entityKindManifest, err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			manifest = newNanifest
		}

*/
